package worker

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/models"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/msgs"
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/utils"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Worker interface {
	DistributeJobs(objs []*models.UploadedVideoFull)
	MasterGenerate()
	VideoInfo()
	SubtitleInfo()
	AudioInfo()
	CreateFolder()
	Upload()
}

func NewWorker(transcoder transcoder.Transcoder, log logger.Logger,
	cfg *config.Config, db storage.StorageI) Worker {
	return &workerPools{
		jobsMap:           sync.Map{},
		folderJobs:        make(chan string, config.JobCount),
		audioJobs:         make(chan string, config.JobCount),
		videoJobs:         make(chan string, config.JobCount),
		subtitleJobs:      make(chan string, config.JobCount),
		masterFileJobs:    make(chan string, config.JobCount),
		objectStorageJobs: make(chan string, config.JobCount),
		opts: opts{
			transcoder: transcoder,
			log:        log,
			cfg:        cfg,
			db:         db,
		},
	}
}

type workerPools struct {
	// jobsMap map[string]struct{} // nolint
	jobsMap sync.Map

	folderJobs        chan string
	audioJobs         chan string
	videoJobs         chan string
	subtitleJobs      chan string
	masterFileJobs    chan string
	objectStorageJobs chan string

	opts opts
}

type opts struct {
	transcoder transcoder.Transcoder
	log        logger.Logger
	cfg        *config.Config
	db         storage.StorageI
}

func (w *workerPools) DistributeJobs(objs []*models.UploadedVideoFull) {
	for _, item := range objs {
		if _, exists := w.jobsMap.Load(item.ID); exists {
			continue
		} else {
			w.jobsMap.Store(item.ID, struct{}{})
		}

		switch item.Stage {
		case config.NewStage:
			w.folderJobs <- item.ID
		case config.AudioStage:
			w.audioJobs <- item.ID
		case config.VideoStage:
			w.videoJobs <- item.ID
		case config.SubtitleStage:
			w.subtitleJobs <- item.ID
		case config.MasterStage:
			w.masterFileJobs <- item.ID
		case config.UploadStage:
			w.objectStorageJobs <- item.ID
		default:
			w.opts.log.Info("NOT FOUND STAGE exiting...")
		}
	}
}

func (w *workerPools) Upload() {
	for job := range w.objectStorageJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.UploadStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		filePath := w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		s3link := w.getS3Link(videoItem.Type, videoItem.MovieSlug, videoItem.SerialSlug)

		err = w.opts.transcoder.UploadToS3(filePath, s3link)
		if err != nil {
			w.opts.log.Error("error while upload", logger.Error(err))
			continue
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) MasterGenerate() {
	for job := range w.masterFileJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.MasterStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		audioList := utils.GetAudioList(videoItem.Audios)
		subtitleList := utils.GetLangArrayString(videoItem.Subtitles)

		var resolutions []ffmpeg.Resolution

		if len(videoItem.Videos) > 0 {
			resolutions = utils.GetResolution(videoItem.Videos[0])
		}

		outputPath := w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		err = w.opts.transcoder.GenerateMasterPlaylist(folder.GenerateMasterOpts{
			AudioList:      audioList,
			SubtitleList:   subtitleList,
			ResolutionList: resolutions,
			InputPath:      inputPath,
			Slug:           videoItem.MovieSlug,
			OutputPath:     outputPath,
		})
		if w.ffmpegError(videoItem.ID, "error while generating master playlist", err) {
			continue
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) VideoInfo() {
out:
	for job := range w.videoJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.VideoStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		outputPath := w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		stream := videoItem.Videos[0]

		resolutionList := utils.GetResolution(stream)

		for _, resolution := range resolutionList {
			err = w.opts.transcoder.ResizeVideo(ffmpeg.ResizeVideoArgs{
				Slug:        videoItem.MovieSlug,
				Input:       inputPath,
				Width:       strconv.Itoa(resolution.Width),
				Height:      strconv.Itoa(resolution.Height),
				BitRate:     resolution.BitRate,
				InputObject: stream,
				OutputPath:  outputPath,
				Disk:        videoItem.Disk,
				Codec:       stream.CodecName,
			})
			if w.ffmpegError(videoItem.ID, "error while extracting video", err) {
				continue out
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.opts.log.Info("video resize done", logger.String("slug", videoItem.MovieSlug))

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) SubtitleInfo() {
out:
	for job := range w.subtitleJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.SubtitleStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		outputPath := w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		for idx, stream := range videoItem.Subtitles {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.opts.transcoder.ExtractSubtitle(inputPath, lang, videoItem.MovieSlug,
				outputPath, videoItem.Disk, idx)

			if w.ffmpegError(videoItem.ID, "error while extracting audio", err) {
				continue out
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.opts.log.Info("subtitle extract done", logger.String("slug", videoItem.MovieSlug))

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) AudioInfo() {
out:
	for job := range w.audioJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.AudioStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		outputPath := w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		for idx, stream := range videoItem.Audios {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.opts.transcoder.ExtractAudio(inputPath, lang, videoItem.MovieSlug, outputPath,
				videoItem.Disk, idx)

			if w.ffmpegError(videoItem.ID, "error while extracting info of audio", err) {
				continue out
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.opts.log.Info("audio extract done", logger.String("slug", videoItem.MovieSlug))

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) CreateFolder() {
	for job := range w.folderJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.NewStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		videos, audios, subtitles, err := w.extractLayers(inputPath)
		if w.ffmpegError(videoItem.ID, "error while extracting info of video", err) {
			continue
		}

		audioList := utils.GetLangArrayString(audios)
		subtitleList := utils.GetLangArrayString(subtitles)
		qualityList := utils.GetVideosArrayString(videos[0])

		var filePath = w.getOutputPath(videoItem.Output, videoItem.MovieSlug)

		err = w.opts.transcoder.GenerateFilesDirectory(folder.FolderOpts{
			OutputPath:   filePath,
			AudioList:    audioList,
			SubtitleList: subtitleList,
			VideoList:    qualityList,
			ScriptPath:   w.opts.cfg.ScriptsFolder,
		})
		if err != nil {
			w.opts.log.Error("error while creating folders", logger.Error(err))
			continue
		}

		err = w.opts.db.UploadedVideo().UpdateStreams(context.Background(), models.UpdateStreams{
			AudioStreams:    audios,
			VideoStreams:    videos,
			SubtitleStreams: subtitles,
			Stage:           config.StagesMatrix[videoItem.Stage],
			ID:              videoItem.ID,
		})
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		w.jobsMap.Delete(job)
	}
}

func (w *workerPools) ffmpegError(id, msg string, err error) bool {
	if err != nil {
		dbErr := w.updateFailure(id, err.Error())
		if dbErr != nil {
			w.opts.log.Error(msgs.ErrUpdFail, logger.Error(err))
		}

		return true
	}

	return false
}

func (w *workerPools) updateFailure(id, msg string) error {
	return w.opts.db.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
		ID:       id,
		Failed:   true,
		ErrorMsg: msg,
	})
}

func (w *workerPools) updateStage(id, stage string) error {
	return w.opts.db.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
		ID:    id,
		Stage: stage,
	})
}

func (w *workerPools) extractLayers(inputPath string) ([]ffmpeg.Stream, []ffmpeg.Stream,
	[]ffmpeg.Stream, error) {
	videos, err := w.opts.transcoder.ExtractInfos(inputPath)
	if err != nil {
		return nil, nil, nil, err
	}

	audios, err := w.opts.transcoder.ExtractAudioLayers(inputPath)
	if err != nil {
		return nil, nil, nil, err
	}

	subtitles, err := w.opts.transcoder.GetSubtitleLayers(inputPath)
	if err != nil {
		return nil, nil, nil, err
	}

	return videos, audios, subtitles, nil
}

// func (w *workerPools) getS3Link(mType, slug string, season, episode int32) string {
// 	if mType == config.Episode {
// 		return fmt.Sprintf("s3://%s/%ss/%s/%s-S%dE%d", w.opts.cfg.BucketName, mType,
// 			slug, slug, season, episode)
// 	}

// 	return fmt.Sprintf("s3://%s/%ss/%s/", w.opts.cfg.BucketName, mType, slug)
// }
func (w *workerPools) getS3Link(mType, slug, epSlug string) string {
	if mType == config.Episode {
		return fmt.Sprintf("s3://%s/%ss/%s/%s", w.opts.cfg.BucketName, "serial", epSlug, slug)
	}

	return fmt.Sprintf("s3://%s/%ss/%s/", w.opts.cfg.BucketName, mType, slug)
}

func (w *workerPools) getOutputPath(path, slug string) string {
	if path == "" {
		return fmt.Sprintf("%s/%s/", w.opts.cfg.OutputDir, slug)
	}

	return fmt.Sprintf("%s/%s/", path, slug)
}
