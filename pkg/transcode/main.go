package transcoder

import "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"

type WorkerPools struct {
	jobMap map[string]struct{}

	InfoJobs          chan string
	FolderJobs        chan string
	AudioJobs         chan string
	VideoJobs         chan string
	SubtitleJobs      chan string
	ObjectStorageJobs chan string
}

type AudioLayerReader interface {
	ExtractInfos()
}

type AuidoExtracter interface {
	ExtractAudios()
	AudioLayerReader
}

type SubtitleLayerReader interface {
	ExtractInfos()
}

type SubtitleExtracter interface {
	SubtitleLayerReader
	ExtractSubtitles()
}

type ObjectUploader interface {
	UploadToS3()
}

type FileFolderGenerator interface {
	GenerateFilesDirecetory()
	GenerateMasterPlaylist()
}

type Transcoder interface {
	AuidoExtracter
	video.VideoExtracter
	SubtitleExtracter
	ObjectUploader
	FileFolderGenerator
}
