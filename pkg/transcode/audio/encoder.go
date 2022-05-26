package audio

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioLayerReader interface {
	ExtractLayers(input string) ([]ffmpeg.Stream, error)
}

type AudioExtracter interface {
	ExtractAudio(input, lang, slug string, index int) error
	AudioLayerReader
}

type AudioObject struct {
	cfg *config.Config
	log logger.Logger

	audio ffmpeg.AudioAPI
}

func NewAudioExtracter(cfg *config.Config, log logger.Logger) AudioExtracter {
	return &AudioObject{
		cfg:   cfg,
		log:   log,
		audio: ffmpeg.NewAudioAPI(cfg, log),
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

func (a *AudioObject) ExtractAudio(input, lang, slug string, index int) error {
	err := a.audio.ExtractAudio(input, lang, slug, index)
	if err != nil {
		a.log.Error("failed to extract audios", logger.Error(err))
		return err
	}

	return nil
}
