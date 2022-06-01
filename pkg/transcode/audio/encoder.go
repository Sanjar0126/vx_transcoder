package audio

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
)

type AudioLayerReader interface {
	ExtractLayers(input string) ([]ffmpeg.Stream, error)
}

type AudioExtracter interface {
	ExtractAudio(input, lang, slug string, index int) error
	AudioLayerReader
}

type AudioObject transcoder.TranscoderStruct

//func NewAudioExtracter(cfg *config.Config, log logger.Logger) AudioExtracter {
//	return &AudioObject{
//		cfg:   cfg,
//		log:   log,
//		audio: ffmpeg.NewAudioAPI(cfg, log),
//	}
//}

func (a *AudioObject) ExtractLayers(input string) ([]ffmpeg.Stream, error) {
	streams, err := a.Audio.GetAudioLayers(input)
	if err != nil {
		a.Log.Error("failed to extract audio layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (a *AudioObject) ExtractAudio(input, lang, slug string, index int) error {
	err := a.Audio.ExtractAudio(input, lang, slug, index)
	if err != nil {
		a.Log.Error("failed to extract audios", logger.Error(err))
		return err
	}

	return nil
}
