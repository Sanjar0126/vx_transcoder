package repo

import (
	"context"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/models"
)

var (
	UploadedVideoCollection = "uploaded_video"
	// UploadedVideoCollection = "uploaded_video2"
)

type UploadedVideoI interface {
	Create(ctx context.Context, req models.UploadedVideoFull) (string, error)
	Get(ctx context.Context, id string) (*models.UploadedVideoFull, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, req models.UploadVideoRequest) error
	UpdateStreams(ctx context.Context, req models.UpdateStreams) error
	GetAll(ctx context.Context,
		filter models.UploadedVideoFilter) ([]*models.UploadedVideoFull, error)
}
