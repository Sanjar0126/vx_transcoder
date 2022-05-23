package ffmpeg

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type ToolFfmpeg struct {
	cfg *config.Config
	log logger.Logger

	video    *VideoAPI
	audio    *AudioAPI
	subtitle *SubtitleAPI
}

func NewToolFfmpeg(cfg *config.Config, log logger.Logger) ToolFfmpeg {
	return ToolFfmpeg{
		cfg: cfg,
		log: log,
	}
}

func (t *ToolFfmpeg) Video() *VideoAPI {
	if t.video == nil {
		t.log.Fatal("video api is not initialized...")
	}

	return t.video
}

func (t *ToolFfmpeg) Audio() *AudioAPI {
	if t.audio == nil {
		t.log.Fatal("audio api is not initialized...")
	}

	return t.audio
}

func (t *ToolFfmpeg) Subtitle() *SubtitleAPI {
	if t.subtitle == nil {
		t.log.Fatal("subtitle api is not initialized...")
	}

	return t.subtitle
}
