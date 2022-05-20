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
		"new":      "info",
		"info":     "audio",
		"audio":    "video",
		"video":    "subtitle",
		"subtitle": "master",
		"master":   "upload",
	}
)
