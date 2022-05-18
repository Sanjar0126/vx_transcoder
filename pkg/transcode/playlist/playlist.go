package playlist

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Playlist struct {
	cfg *config.Config
	log logger.Logger
	db  storage.StorageI
}

func NewPlaylist(cfg *config.Config, log logger.Logger, db storage.StorageI) interfaces.PlaylistI {
	return &Playlist{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (p *Playlist) GenerateMasterPlaylist(input string) error {
	fmt.Println("Generating master playlist", input)

	return nil
}
