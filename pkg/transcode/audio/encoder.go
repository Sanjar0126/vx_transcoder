package audio

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioLayerReader interface {
	ExtractLayers(input string) ([]ffmpeg.Stream, error)
}

type AuidoExtracter interface {
	ExtractAudio(input, lang string, inputObject ffmpeg.Stream) error
	AudioLayerReader
}

type AudioObject struct {
	cfg *config.Config
	log logger.Logger

	audio ffmpeg.AudioAPI
}

func NewAudioExtracter(cfg *config.Config, log logger.Logger) AuidoExtracter {
	return &AudioObject{
		cfg: cfg,
		log: log,
	}
}

func (a *AudioObject) ExtractLayers(input string) ([]ffmpeg.Stream, error) {
	streams, err := a.audio.GetAudioLayers(input)
	if err != nil {
		a.log.Error("failed to extract audio layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (a *AudioObject) ExtractAudio(input, lang string, inputObject ffmpeg.Stream) error {
	err := a.audio.ExtractAudio(input, lang, inputObject)
	if err != nil {
		a.log.Error("failed to exract audios", logger.Error(err))
		return err
	}

	return nil
}
