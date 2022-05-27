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
		"new":      "audio",
		"audio":    "subtitle",
		"subtitle": "video",
		"video":    "master",
		"master":   "upload",
	}

	StagesArray = []string{"new", "audio", "video", "subtitle", "master"}

	JobCount = 3

	InputPathTemplate = "%s%s.%s"
)
