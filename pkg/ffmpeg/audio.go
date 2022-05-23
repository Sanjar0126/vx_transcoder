package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioAPI struct {
	cfg *config.Config
	log logger.Logger

	codecType string
}

func NewAudioAPI(cfg *config.Config, log logger.Logger) AudioAPI {
	return AudioAPI{
		cfg:       cfg,
		log:       log,
		codecType: "audio",
	}
}

func (a *AudioAPI) GetAudioLayers(input string) ([]Stream, error) {
	audioLayers, err := getLayers(a.cfg.ScriptsFolder, input, a.codecType)
	if err != nil {
		a.log.Error("failed to retrieve audio layers", logger.Error(err))
		return nil, err
	}

	a.log.Info("audio layers are retreived", logger.Any("audios", audioLayers))

	return audioLayers, nil
}

func (a *AudioAPI) ExtractAudio() {

}
