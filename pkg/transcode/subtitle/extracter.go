package subtitle

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type SubtitleLayerReader interface {
	GetSubtitleLayers(input string) ([]ffmpeg.Stream, error)
}

type SubtitleExtracter interface {
	SubtitleLayerReader
	ExtractSubtitle(input, lang string, inputObject ffmpeg.Stream) error
}

type SubtitleObject struct {
	cfg *config.Config
	log logger.Logger

	subtitle ffmpeg.SubtitleAPI
}

func NewSubtitleExtracter(cfg *config.Config, log logger.Logger) SubtitleExtracter {
	return &SubtitleObject{
		cfg: cfg,
		log: log,
	}
}

func (s *SubtitleObject) GetSubtitleLayers(input string) ([]ffmpeg.Stream, error) {
	streams, err := s.subtitle.GetSubtitleLayers(input)
	if err != nil {
		s.log.Error("failed to get subtitle layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (s *SubtitleObject) ExtractSubtitle(input, lang string, inputObject ffmpeg.Stream) error {
	err := s.subtitle.ExtractSubtitle(input, lang, inputObject)
	if err != nil {
		s.log.Error("failed to extract subtitles", logger.Error(err))
		return err
	}

	return nil
}
