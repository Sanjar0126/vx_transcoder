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
	ExtractSubtitle(input, lang, slug, output string, index int) error
}

type SubtitleObject struct {
	cfg *config.Config
	log logger.Logger

	subtitle ffmpeg.SubtitleAPI
}

func NewSubtitleExtracter(cfg *config.Config, log logger.Logger) SubtitleExtracter {
	return &SubtitleObject{
		cfg:      cfg,
		log:      log,
		subtitle: ffmpeg.NewSubtitleAPI(cfg, log),
	}
}

func (s *SubtitleObject) GetSubtitleLayers(input string) ([]ffmpeg.Stream, error) {
	return s.subtitle.GetSubtitleLayers(input)
}

func (s *SubtitleObject) ExtractSubtitle(input, lang, slug, output string, index int) error {
	return s.subtitle.ExtractSubtitle(input, lang, slug, output, index)
}
