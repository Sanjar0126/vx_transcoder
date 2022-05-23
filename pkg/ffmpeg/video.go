package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoAPI struct {
	cfg *config.Config
	log logger.Logger
}

func NewVideoAPI(cfg *config.Config, log logger.Logger) VideoAPI {
	return VideoAPI{
		cfg: cfg,
		log: log,
	}
}

func (v *VideoAPI) GetVideoWidthAndHeight() {

}

func (v *VideoAPI) GetVideoLayers() {

}

func (v *VideoAPI) GetVideoDuration() {

}

func (v *VideoAPI) ConvertVideo() {

}
