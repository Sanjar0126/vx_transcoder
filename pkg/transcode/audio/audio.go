package audio

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Audio struct {
	cfg *config.Config
	log logger.Logger
	db  storage.StorageI
}

func NewAudio(cfg *config.Config, log logger.Logger, db storage.StorageI) interfaces.AudioI {
	return &Audio{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (a *Audio) RetreiveAudioChannels(input string) string {
	fmt.Println("Retreiving audio channels", input)

	return ""
}

func (a *Audio) ExtractAudio(input string) error {
	fmt.Println("Extracting audio")

	return nil
}
