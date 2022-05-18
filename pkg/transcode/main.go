package transcode

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/disk"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/playlist"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type TranscoderI interface {
	Audio() interfaces.AudioI
	Video() interfaces.VideoI
	Subtitle() interfaces.SubtitleI
	Playlist() interfaces.PlaylistI
	Disk() interfaces.DiskI
	Run()
}

func New(cfg *config.Config, log logger.Logger, db storage.StorageI) TranscoderI {
	return &transcoder{
		audio:    audio.NewAudio(cfg, log, db),
		video:    video.NewVideo(cfg, log, db),
		subtitle: subtitle.NewSubtitle(cfg, log, db),
		playlist: playlist.NewPlaylist(cfg, log, db),
		disk:     disk.NewDisk(cfg, log, db),
		db:       db,
	}
}

func (t *transcoder) Audio() interfaces.AudioI {
	return t.audio
}

func (t *transcoder) Video() interfaces.VideoI {
	return t.video
}

func (t *transcoder) Subtitle() interfaces.SubtitleI {
	return t.subtitle
}

func (t *transcoder) Playlist() interfaces.PlaylistI {
	return t.playlist
}

func (t *transcoder) Disk() interfaces.DiskI {
	return t.disk
}
