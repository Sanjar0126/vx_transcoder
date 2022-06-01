package video

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoObject struct {
	cfg   *config.Config
	log   logger.Logger
	video ffmpeg.VideoAPI
}

type VideoLayerReader interface {
	ExtractInfos(input string) ([]ffmpeg.Stream, error)
}

type VideoExtracter interface {
	VideoLayerReader
	ResizeVideo(args ffmpeg.ResizeVideoArgs) error
}

func NewVideoExtractor(cfg *config.Config, log logger.Logger) VideoExtracter {
	return &VideoObject{
		cfg:   cfg,
		log:   log,
		video: ffmpeg.NewVideoAPI(cfg, log),
	}
}

func (v *VideoObject) ExtractInfos(input string) ([]ffmpeg.Stream, error) {
	return v.video.GetVideoLayers(input)
}

func (v *VideoObject) ResizeVideo(args ffmpeg.ResizeVideoArgs) error {
	return v.video.ResizeVideo(args)
}
