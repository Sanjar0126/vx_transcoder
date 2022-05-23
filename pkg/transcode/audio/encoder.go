package audio

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioLayerReader interface {
	ExtractInfos()
}

type AuidoExtracter interface {
	ExtractAudios()
	AudioLayerReader
}

type AudioObject struct {
	cfg *config.Config
	log logger.Logger
}

func NewAudioExtracter(cfg *config.Config, log logger.Logger) AuidoExtracter {
	return &AudioObject{
		cfg: cfg,
		log: log,
	}
}

func (a *AudioObject) ExtractInfos() {
	panic("not implemented yet")
}

func (a *AudioObject) ExtractAudios() {
	panic("not implemented yet")
}
