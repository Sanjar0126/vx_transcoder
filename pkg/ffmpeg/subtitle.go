package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type SubtitleAPI struct {
	cfg *config.Config
	log logger.Logger
}

func NewSubtitleAPI(cfg *config.Config, log logger.Logger) SubtitleAPI {
	return SubtitleAPI{
		cfg: cfg,
		log: log,
	}
}

func (s *SubtitleAPI) ExtractSubtitle() {

}

func (s *SubtitleAPI) GetSubtitleLayers() {

}
