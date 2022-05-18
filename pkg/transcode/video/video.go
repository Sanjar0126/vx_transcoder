package video

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Video struct {
	cfg *config.Config
	log logger.Logger
	db  storage.StorageI
}

func NewVideo(cfg *config.Config, log logger.Logger, db storage.StorageI) interfaces.VideoI {
	return &Video{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (v *Video) GetVideoDuration(input string) (int, error) {
	fmt.Println("Extracting video duration", input)

	return 0, nil
}

func (v *Video) GetVideoResolution(input string) error {
	fmt.Println("Extracting video resolution", input)

	return nil
}

func (v *Video) GetVideoWidthHeight(input string) error {
	fmt.Println("Extracting video width and height", input)

	return nil
}

func (v *Video) Transcode(input string) error {
	fmt.Println("Transcoding", input)

	return nil
}

func (v *Video) TranscodeVideoByResolution(input string) error {
	fmt.Println("Transcoding by resolution", input)

	return nil
}
