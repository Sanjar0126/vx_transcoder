package transcoder

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

type VideoLayerReader interface {
	ExtractInfos()
	ExtractDurations()
}

type VideoExtracter interface {
	VideoLayerReader
	ConvertVideos(inputType, outputType string)
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
	VideoExtracter
	SubtitleExtracter
	ObjectUploader
	FileFolderGenerator
}
