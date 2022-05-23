package cloud

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type ObjectUploader interface {
	UploadToS3()
}

type ObjectStorage struct {
	cfg *config.Config
	log logger.Logger
}

func NewObjectUploader(cfg *config.Config, log logger.Logger) ObjectUploader {
	return &ObjectStorage{
		cfg: cfg,
		log: log,
	}
}

func (o *ObjectStorage) UploadToS3() {

}
