package subtitle

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type SubtitleLayerReader interface {
	ExtractInfos()
}

type SubtitleExtracter interface {
	SubtitleLayerReader
	ExtractSubtitles()
}

type SubtitleObject struct {
	cfg *config.Config
	log logger.Logger
}

func NewSubtitleExtracter(cfg *config.Config, log logger.Logger) SubtitleExtracter {
	return &SubtitleObject{
		cfg: cfg,
		log: log,
	}
}

func (s *SubtitleObject) ExtractInfos() {

}

func (s *SubtitleObject) ExtractSubtitles() {

}
