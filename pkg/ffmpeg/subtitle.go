package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type SubtitleAPI struct {
	cfg       *config.Config
	log       logger.Logger
	codecType string
}

func NewSubtitleAPI(cfg *config.Config, log logger.Logger) SubtitleAPI {
	return SubtitleAPI{
		cfg:       cfg,
		log:       log,
		codecType: "subtitle",
	}
}

func (s *SubtitleAPI) GetSubtitleLayers(input string) ([]Streams, error) {
	subtitleLayers, err := getLayers(s.cfg.ScriptsFolder, input, s.codecType)
	if err != nil {
		s.log.Error("failed to retrieve subtitle layers", logger.Error(err))
		return nil, err
	}

	s.log.Info("subtitle layers are retreived", logger.Any("subtitles", subtitleLayers))

	return subtitleLayers, nil
}

func (s *SubtitleAPI) ExtractSubtitle(input string) {

}
