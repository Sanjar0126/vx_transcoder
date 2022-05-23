package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioAPI struct {
	cfg *config.Config
	log logger.Logger
}

func NewAudioAPI(cfg *config.Config, log logger.Logger) AudioAPI {
	return AudioAPI{
		cfg: cfg,
		log: log,
	}
}

func (a *AudioAPI) GetAudioLayers() {

}

func (a *AudioAPI) ExtractAudio() {

}
