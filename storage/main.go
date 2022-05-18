package storage

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage/mongo"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage/repo"
	db "go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	UploadedVideo() repo.UploadedVideoI
}

type storageMDB struct {
	uploadedVideoRepo repo.UploadedVideoI
}

func New(db *db.Database) StorageI {
	return &storageMDB{
		uploadedVideoRepo: mongo.NewUploadedVideoRepo(db),
	}
}

func (s *storageMDB) UploadedVideo() repo.UploadedVideoI {
	return s.uploadedVideoRepo
}
