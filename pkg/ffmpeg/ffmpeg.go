package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type ToolFfmpeg struct {
	cfg *config.Config
	log logger.Logger
}

func NewToolFfmpeg(cfg *config.Config, log logger.Logger) ToolFfmpeg {
	return ToolFfmpeg{
		cfg: cfg,
		log: log,
	}
}

func (t *ToolFfmpeg) GetVideoWidthAndHeight() {

}

func (t *ToolFfmpeg) GetVideoLayers() {

}

func (t *ToolFfmpeg) GetAudioLayers() {

}

func (t *ToolFfmpeg) GetSubtitleLayers() {

}

func (t *ToolFfmpeg) GetVideoDuration() {

}

func (t *ToolFfmpeg) ConvertVideo() {

}

func (t *ToolFfmpeg) ExtractAudio() {

}

func (t *ToolFfmpeg) ExtractSubtitle() {

}
