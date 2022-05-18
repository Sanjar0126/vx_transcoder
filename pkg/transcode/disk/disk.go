package disk

import (
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Disk struct {
	cfg *config.Config
	log logger.Logger
	db  storage.StorageI
}

func NewDisk(cfg *config.Config, log logger.Logger, db storage.StorageI) interfaces.DiskI {
	return &Disk{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (d *Disk) CreateFolder(input string) error {
	fmt.Println("Creating folder", input)

	return nil
}

func (d *Disk) UploadToCloud(input string) error {
	fmt.Println("Upload to cloud", input)

	return nil
}

func (d *Disk) MakeDirectoryPath(key string) string {
	inputPath := fmt.Sprintf("%s/%s", d.cfg.InputDirectory, key)

	return inputPath
}

func (d *Disk) MakeCdnPath(key string) string {
	inputPath := fmt.Sprintf("https://voxe-temp-1.s3.eu-north-1.amazonaws.com/%v", key)

	return inputPath
}
