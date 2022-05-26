package config

var (
	EnvDevelop             = "develop"
	SlugLength         int = 30
	Langs                  = []string{"ru", "en", "uz"}
	LangUz                 = "uz"
	LangRu                 = "ru"
	LangEn                 = "en"
	ContentServicePort     = 9100
	RabbitMQPort           = 5672

	StagesMatrix = map[string]string{
		"info":     "audio",
		"new":      "audio",
		"audio":    "subtitle",
		"subtitle": "video",
		"video":    "master",
		"master":   "upload",
	}

	StagesArray = []string{"new", "info", "audio", "video", "subtitle", "master"}

	JobCount = 2
)
