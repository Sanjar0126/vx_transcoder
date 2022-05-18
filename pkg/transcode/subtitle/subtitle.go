package subtitle

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Subtitle struct {
	cfg *config.Config
	log logger.Logger
	db  storage.StorageI
}

func NewSubtitle(cfg *config.Config, log logger.Logger, db storage.StorageI) interfaces.SubtitleI {
	return &Subtitle{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (s *Subtitle) RetreiveSubtitleChannels(input string) string {
	fmt.Println("Retreiving subtitle channels", input)

	return ""
}

func (s *Subtitle) ExtractSubtitle(input string) error {
	fmt.Println("Extracting subtitle channels", input)

	return nil
}
