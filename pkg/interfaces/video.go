package interfaces

type VideoI interface {
	GetVideoDuration(input string) (int, error)
	GetVideoResolution(input string) error
	GetVideoWidthHeight(input string) error
	Transcode(input string) error
	TranscodeVideoByResolution(input string) error
}
