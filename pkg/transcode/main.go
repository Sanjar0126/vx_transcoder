package transcoder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/cloud"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
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
	FolderGen folder.FileFolderGenerator
	VideoExt  video.VideoExtracter
	AudioExt  audio.AudioExtracter
	SubExt    subtitle.SubtitleExtracter

	Log logger.Logger
	Cfg *config.Config
	DB  storage.StorageI
}

type Transcoder interface {
	audio.AudioExtracter
	video.VideoExtracter
	subtitle.SubtitleExtracter
	folder.FileFolderGenerator
	cloud.ObjectUploader
}
