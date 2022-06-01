package cloud

import (
	"errors"
	"fmt"
	"os/exec"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type ObjectUploader interface {
	UploadToS3(input, out string) error
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

func (o *ObjectStorage) UploadToS3(input, out string) error {
	var (
		script = fmt.Sprintf("%s%s", o.cfg.ScriptsFolder, "/upload_s3.sh")
	)

	o.log.Info("uploading file info", logger.String("input", input))

	_, err := exec.Command(
		"/bin/sh",
		script,
		out,
	).Output()

	if err != nil {
		o.log.Error("failed to upload file", logger.Error(err))
		return errors.New("failed to upload file")
	}

	return nil
}
