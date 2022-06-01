package video

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
)

type VideoObject transcoder.TranscoderStruct

type VideoLayerReader interface {
	ExtractInfos(input string) ([]ffmpeg.Stream, error)
}

type VideoExtracter interface {
	VideoLayerReader
	ResizeVideo(args ffmpeg.ResizeVideoArgs) error
}

//func NewVideoExtractor(cfg *config.Config, log logger.Logger) VideoExtracter {
//	return &VideoObject{
//		cfg:   cfg,
//		log:   log,
//		video: ffmpeg.NewVideoAPI(cfg, log),
//	}
//}

func (v *VideoObject) ExtractInfos(input string) ([]ffmpeg.Stream, error) {
	streams, err := v.Video.GetVideoLayers(input)
	if err != nil {
		v.Log.Info("failed to retrieve info about video layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (v *VideoObject) ResizeVideo(args ffmpeg.ResizeVideoArgs) error {
	err := v.Video.ResizeVideo(args)
	if err != nil {
		v.Log.Error("failed to resize video", logger.Error(err))
		return err
	}

	return nil
}
