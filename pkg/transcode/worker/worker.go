package worker

import (
	"context"
	"fmt"
	"strconv"

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
	workerPools := &workerPools{
		jobsMap:           make(map[string]struct{}),
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

	initializeGoroutines(workerPools)

	return workerPools
}

type workerPools struct {
	jobsMap map[string]struct{} // nolint

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

func initializeGoroutines(workerPool *workerPools) {
	go workerPool.CreateFolder()
	go workerPool.AudioInfo()
	go workerPool.SubtitleInfo()
	go workerPool.VideoInfo()
	go workerPool.MasterGenerate()
	go workerPool.Upload()
}

func (w *workerPools) DistributeJobs(objs []*models.UploadedVideoFull) {
	for _, item := range objs {
		if _, exists := w.jobsMap[item.ID]; exists {
			w.opts.log.Info("already exists in queue")
		} else {
			w.jobsMap[item.ID] = struct{}{}
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
	for job := range w.masterFileJobs {
		videoItem, err := w.opts.db.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.opts.log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		if videoItem.Stage != config.UploadStage {
			w.opts.log.Error(msgs.WrongStage)
			continue
		}

		filePath := w.opts.cfg.OutputDir + "/" + videoItem.MovieSlug
		s3link := fmt.Sprintf("s3://%s/%ss/%s", w.opts.cfg.BucketName,
			videoItem.Type, videoItem.MovieSlug)

		err = w.opts.transcoder.UploadToS3(filePath, s3link)
		if err != nil {
			w.opts.log.Error("error while upload", logger.Error(err))
			continue
		}

		delete(w.jobsMap, job)
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

		err = w.opts.transcoder.GenerateMasterPlaylist(folder.GenerateMasterOpts{
			AudioList:      audioList,
			SubtitleList:   subtitleList,
			ResolutionList: resolutions,
			InputPath:      inputPath,
			Slug:           videoItem.MovieSlug,
		})
		if err != nil {
			w.opts.log.Error("error while generating master playlist", logger.Error(err))
			continue
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.jobsMap, job)
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

		for _, stream := range videoItem.Videos {
			resolutionList := utils.GetResolution(stream)

			for _, resolution := range resolutionList {
				err = w.opts.transcoder.ResizeVideo(ffmpeg.ResizeVideoArgs{
					Slug:        videoItem.MovieSlug,
					Input:       inputPath,
					Width:       strconv.Itoa(resolution.Width),
					Height:      strconv.Itoa(resolution.Height),
					BitRate:     resolution.BitRate,
					InputObject: stream,
				})
				if err != nil {
					w.opts.log.Error("error while extracting video", logger.Error(err))
					continue out
				}
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.jobsMap, job)
	}
}

func (w *workerPools) SubtitleInfo() {
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

		for idx, stream := range videoItem.Subtitles {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.opts.transcoder.ExtractSubtitle(inputPath, lang, videoItem.MovieSlug, idx)

			if err != nil {
				w.opts.log.Error("error while extracting audio", logger.Error(err))
				continue
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.jobsMap, job)
	}
}

func (w *workerPools) AudioInfo() {
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

		for idx, stream := range videoItem.Audios {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.opts.transcoder.ExtractAudio(inputPath, lang, videoItem.MovieSlug, idx)

			if err != nil {
				w.opts.log.Error("error while extracting audio", logger.Error(err))
				continue
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.opts.log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.jobsMap, job)
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
		if err != nil {
			w.opts.log.Error("error while extracting info of video", logger.Error(err))
			continue
		}

		audioList := utils.GetLangArrayString(audios)
		subtitleList := utils.GetLangArrayString(subtitles)
		qualityList := utils.GetVideosArrayString(videos)

		err = w.opts.transcoder.GenerateFilesDirectory(folder.FolderOpts{
			OutputPath:   w.opts.cfg.OutputDir + "/" + videoItem.MovieSlug,
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

		delete(w.jobsMap, job)
	}
}

func (w *workerPools) updateStage(id, stage string) error {
	err := w.opts.db.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
		ID:    id,
		Stage: stage,
	})

	return err
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
