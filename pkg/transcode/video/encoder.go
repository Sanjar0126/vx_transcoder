package video

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoObject struct {
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

func NewVideoExtracter(cfg *config.Config, log logger.Logger) VideoExtracter {
	return &VideoObject{
		cfg: cfg,
		log: log,
	}
}

func (v *VideoObject) ExtractInfos() {
	panic("not implemented yet")
}

func (v *VideoObject) ExtractDurations() {
	panic("not implemented yet")
}

func (v *VideoObject) ConvertVideos(inputType, outputType string) {
	panic("not implemented yet")
}
