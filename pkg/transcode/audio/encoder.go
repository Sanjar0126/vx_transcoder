package audio

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioLayerReader interface {
	ExtractAudioLayers(input string) ([]ffmpeg.Stream, error)
}

type AudioExtractor interface {
	ExtractAudio(input, lang, slug string, index int) error
	AudioLayerReader
}

type AudioObject struct {
	cfg *config.Config
	log logger.Logger

	audio ffmpeg.AudioAPI
}

func NewAudioExtractor(cfg *config.Config, log logger.Logger) AudioExtractor {
	return &AudioObject{
		cfg:   cfg,
		log:   log,
		audio: ffmpeg.NewAudioAPI(cfg, log),
	}
}

func (a *AudioObject) ExtractAudioLayers(input string) ([]ffmpeg.Stream, error) {
	return a.audio.GetAudioLayers(input)
}

func (a *AudioObject) ExtractAudio(input, lang, slug string, index int) error {
	return a.audio.ExtractAudio(input, lang, slug, index)
}
