package cloud

import (
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
)

type ObjectUploader interface {
	UploadToS3()
}

type ObjectStorage transcoder.TranscoderStruct

//func NewObjectUploader(cfg *config.Config, log logger.Logger) ObjectUploader {
//	return &ObjectStorage{
//		cfg: cfg,
//		log: log,
//	}
//}

func (o *ObjectStorage) UploadToS3() {

}
