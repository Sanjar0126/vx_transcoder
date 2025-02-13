package config

var (
	EnvDevelop         = "develop"
	SlugLength         = 30
	Langs              = []string{"ru", "en", "uz"}
	LangUz             = "uz"
	LangRu             = "ru"
	LangEn             = "en"
	ContentServicePort = 9100
	RabbitMQPort       = 5672

	StagesMatrix = map[string]string{
		"new":      "audio",
		"audio":    "subtitle",
		"subtitle": "video",
		"video":    "master",
		"master":   "upload",
		"upload":   "finished",
	}

	StagesArray = []string{"new", "audio", "video", "subtitle", "master"}

	JobCount            = 4
	VideoResizeJobCount = 4
	UploadJobCount      = 4

	InputPathTemplate = "%s/%s.%s"

	NewStage      = "new"
	AudioStage    = "audio"
	SubtitleStage = "subtitle"
	VideoStage    = "video"
	MasterStage   = "master"
	UploadStage   = "upload"
	FinishedStage = "finished"

	Episode = "episode"

	DiskArray = []string{"Expansion", "H1", "T2"}
)
