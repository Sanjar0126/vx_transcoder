package transcoder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/cloud"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
)

type Transcoder interface {
	audio.AudioExtracter
	video.VideoExtracter
	subtitle.SubtitleExtracter
	folder.FileFolderGenerator
	cloud.ObjectUploader
}

type TranscoderStruct struct {
	Cfg *config.Config
	Log logger.Logger

	Video    ffmpeg.VideoAPI
	Audio    ffmpeg.AudioAPI
	Subtitle ffmpeg.SubtitleAPI
}

func NewTranscoder(cfg *config.Config, log logger.Logger) Transcoder {
	return &TranscoderStruct{
		Cfg:      cfg,
		Log:      log,
		Video:    ffmpeg.NewVideoAPI(cfg, log),
		Audio:    ffmpeg.NewAudioAPI(cfg, log),
		Subtitle: ffmpeg.NewSubtitleAPI(cfg, log),
	}
}
