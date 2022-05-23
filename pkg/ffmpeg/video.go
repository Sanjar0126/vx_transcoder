package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoAPI struct {
	cfg       *config.Config
	log       logger.Logger
	codecType string
}

func NewVideoAPI(cfg *config.Config, log logger.Logger) VideoAPI {
	return VideoAPI{
		cfg:       cfg,
		log:       log,
		codecType: "video",
	}
}

func (v *VideoAPI) GetVideoWidthAndHeight() {

}

func (v *VideoAPI) GetVideoLayers(input string) ([]Streams, error) {
	videoLayers, err := getLayers(v.cfg.ScriptsFolder, input, v.codecType)
	if err != nil {
		v.log.Error("failed to retrieve video layers", logger.Error(err))
		return nil, err
	}

	v.log.Info("video layers are retreived", logger.Any("videos", videoLayers))

	return videoLayers, nil
}

func (v *VideoAPI) GetVideoDuration() {

}

func (v *VideoAPI) ConvertVideo() {

}
