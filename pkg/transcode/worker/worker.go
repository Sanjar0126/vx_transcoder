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
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/utils"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type WorkerPools struct {
	JobsMap map[string]struct{} // nolint

	FolderJobs        chan string
	AudioJobs         chan string
	VideoJobs         chan string
	SubtitleJobs      chan string
	MasterInfoJobs    chan string
	ObjectStorageJobs chan string

	Opts Opts
}

type Opts struct {
	Generator folder.FileFolderGenerator
	VideoExt  video.VideoExtracter
	AudioExt  audio.AudioExtracter
	SubExt    subtitle.SubtitleExtracter

	Log logger.Logger
	Cfg *config.Config
	DB  storage.StorageI
}

func (w *WorkerPools) DistributeJobs(objs []*models.UploadedVideoFull) {
	for _, item := range objs {
		if _, exists := w.JobsMap[item.ID]; exists {
			fmt.Println("already exists in queue")
		} else {
			// fmt.Println("new job arrived...")
			w.JobsMap[item.ID] = struct{}{}
		}

		fmt.Println("INPUT ITEM STAGE: ", item.Stage)

		switch item.Stage {
		case "new":
			w.FolderJobs <- item.ID
		case "audio":
			w.AudioJobs <- item.ID
		case "video":
			w.VideoJobs <- item.ID
		case "subtitle":
			w.SubtitleJobs <- item.ID
		case "master":
			w.MasterInfoJobs <- item.ID
		default:
			fmt.Println("NOT FOUND STAGE exiting")
			w.Opts.Log.Info("NOT FOUND STAGE exiting...")
		}
	}
}

func (w *WorkerPools) MasterGenerate() {
	for job := range w.MasterInfoJobs {
		videoItem, err := w.Opts.DB.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.Opts.Log.Error(msgs.ErrDBGetAll, logger.Error(err))
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

		err = w.Opts.Generator.GenerateMasterPlaylist(folder.GenerateMasterOpts{
			AudioList:      audioList,
			SubtitleList:   subtitleList,
			ResolutionList: resolutions,
			InputPath:      inputPath,
			Slug:           videoItem.MovieSlug,
		})
		if err != nil {
			w.Opts.Log.Error("error while generating master playlist", logger.Error(err))
			continue
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.Opts.Log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.JobsMap, job)
	}
}

func (w *WorkerPools) VideoInfo() {
out:
	for job := range w.VideoJobs {
		videoItem, err := w.Opts.DB.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.Opts.Log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		for _, stream := range videoItem.Videos {
			resolutionList := utils.GetResolution(stream)

			for _, resolution := range resolutionList {
				err = w.Opts.VideoExt.ResizeVideo(ffmpeg.ResizeVideoArgs{
					Slug:        videoItem.MovieSlug,
					Input:       inputPath,
					Width:       strconv.Itoa(resolution.Width),
					Height:      strconv.Itoa(resolution.Height),
					BitRate:     resolution.BitRate,
					InputObject: stream,
				})
				if err != nil {
					w.Opts.Log.Error("error while extracting video", logger.Error(err))
					continue out
				}
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.Opts.Log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.JobsMap, job)
	}
}

func (w *WorkerPools) SubtitleInfo() {
	for job := range w.SubtitleJobs {
		videoItem, err := w.Opts.DB.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.Opts.Log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		for idx, stream := range videoItem.Subtitles {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.Opts.SubExt.ExtractSubtitle(inputPath, lang, videoItem.MovieSlug, idx)

			if err != nil {
				w.Opts.Log.Error("error while extracting audio", logger.Error(err))
				continue
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.Opts.Log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.JobsMap, job)
	}
}

func (w *WorkerPools) AudioInfo() {
	for job := range w.AudioJobs {
		videoItem, err := w.Opts.DB.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.Opts.Log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		for idx, stream := range videoItem.Audios {
			lang := utils.GetTag(stream.Tags, idx).Language
			err = w.Opts.AudioExt.ExtractAudio(inputPath, lang, videoItem.MovieSlug, idx)

			if err != nil {
				w.Opts.Log.Error("error while extracting audio", logger.Error(err))
				continue
			}
		}

		err = w.updateStage(videoItem.ID, config.StagesMatrix[videoItem.Stage])
		if err != nil {
			w.Opts.Log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.JobsMap, job)
	}
}

func (w *WorkerPools) CreateFolder() {
	for job := range w.FolderJobs {
		videoItem, err := w.Opts.DB.UploadedVideo().Get(context.Background(), job)
		if err != nil {
			w.Opts.Log.Error(msgs.ErrDBGetAll, logger.Error(err))
			continue
		}

		inputPath := fmt.Sprintf(config.InputPathTemplate, videoItem.Path, videoItem.MovieSlug,
			videoItem.Extension)

		videos, err := w.Opts.VideoExt.ExtractInfos(inputPath)
		if err != nil {
			w.Opts.Log.Error("error while extracting info of video", logger.Error(err))
			continue
		}

		audios, err := w.Opts.AudioExt.ExtractLayers(inputPath)
		if err != nil {
			w.Opts.Log.Error("error while extracting audio layers", logger.Error(err))
			continue
		}

		subtitles, err := w.Opts.SubExt.GetSubtitleLayers(inputPath)
		if err != nil {
			w.Opts.Log.Error("error while extracting subtitle layers", logger.Error(err))
			continue
		}

		audioList := utils.GetLangArrayString(audios)
		subtitleList := utils.GetLangArrayString(subtitles)
		qualityList := utils.GetVideosArrayString(videos)

		err = w.Opts.Generator.GenerateFilesDirectory(folder.FolderOpts{
			OutputPath:   w.Opts.Cfg.OutputDir + "/" + videoItem.MovieSlug,
			AudioList:    audioList,
			SubtitleList: subtitleList,
			VideoList:    qualityList,
			ScriptPath:   w.Opts.Cfg.ScriptsFolder,
		})
		if err != nil {
			w.Opts.Log.Error("error while creating folders", logger.Error(err))
			continue
		}

		err = w.Opts.DB.UploadedVideo().UpdateStreams(context.Background(), models.UpdateStreams{
			AudioStreams:    audios,
			VideoStreams:    videos,
			SubtitleStreams: subtitles,
			Stage:           config.StagesMatrix[videoItem.Stage],
			ID:              videoItem.ID,
		})
		if err != nil {
			w.Opts.Log.Error(msgs.ErrUpdStage, logger.Error(err))
			return
		}

		delete(w.JobsMap, job)
	}
}

func (w *WorkerPools) updateStage(id, stage string) error {
	fmt.Println("stage: ", stage)
	err := w.Opts.DB.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
		ID:    id,
		Stage: stage,
	})

	return err
}
