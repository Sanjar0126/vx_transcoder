package transcoder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
)

type WorkerPools struct {
	jobMap map[string]struct{}

	InfoJobs          chan string
	FolderJobs        chan string
	AudioJobs         chan string
	VideoJobs         chan string
	SubtitleJobs      chan string
	ObjectStorageJobs chan string
}

type ObjectUploader interface {
	UploadToS3()
}

type FileFolderGenerator interface {
	GenerateFilesDirecetory()
	GenerateMasterPlaylist()
}

type Transcoder interface {
	audio.AuidoExtracter
	video.VideoExtracter
	subtitle.SubtitleExtracter
	ObjectUploader
	FileFolderGenerator
}
