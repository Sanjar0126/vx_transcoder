package video

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoEncoder struct {
	cfg *config.Config
	log logger.Logger
}

type VideoLayerReader interface {
	ExtractInfos()
	ExtractDurations()
}

type VideoExtracter interface {
	VideoLayerReader
	ConvertVideos(inputType, outputType string)
}

func NewVideoEncoder(cfg *config.Config, log logger.Logger) VideoExtracter {
	return &VideoEncoder{
		cfg: cfg,
		log: log,
	}
}

func (v *VideoEncoder) ExtractInfos() {
	panic("not implemented yet")
}

func (v *VideoEncoder) ExtractDurations() {
	panic("not implemented yet")
}

func (v *VideoEncoder) ConvertVideos(inputType, outputType string) {
	panic("not implemented yet")
}
