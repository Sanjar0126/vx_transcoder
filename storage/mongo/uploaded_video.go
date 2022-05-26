package mongo

import (
	"context"
	"time"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/models"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type uploadedVideoStorage struct {
	collection *mongo.Collection
}

func NewUploadedVideoRepo(db *mongo.Database) repo.UploadedVideoI {
	uploadedVideo := uploadedVideoStorage{
		collection: db.Collection(repo.UploadedVideoCollection),
	}

	return &uploadedVideo
}

func (up *uploadedVideoStorage) Create(ctx context.Context, req models.UploadedVideoFull) (string,
	error) {
	var (
		resp models.UploadedVideo
	)

	_, err := up.collection.InsertOne(ctx, models.UploadedVideoFull{
		//ID:            req.ID,
		MovieSlug:     req.MovieSlug,
		Type:          req.Type,
		Stage:         req.Stage,
		SeasonNumber:  req.SeasonNumber,
		EpisodeNumber: req.EpisodeNumber,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})

	if err != nil {
		return resp.FilePath, err
	}

	resp.FilePath = req.ID

	return resp.FilePath, nil
}

func (up *uploadedVideoStorage) Update(ctx context.Context, req models.UploadVideoRequest) error {
	objId, _ := primitive.ObjectIDFromHex(req.ID)

	err := up.collection.FindOneAndUpdate(ctx, bson.M{"_id": objId},
		bson.M{"$set": bson.M{"stage": req.Stage}})

	if err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (up *uploadedVideoStorage) GetAll(ctx context.Context,
	filter models.UploadedVideoFilter) ([]*models.UploadedVideoFull, error) {
	var (
		response []*models.UploadedVideoFull
	)

	rows, err := up.collection.Find(ctx, bson.M{
		"stage": bson.M{
			"$in": filter.Stages,
		},
	})
	if err != nil {
		return response, err
	}

	if err := rows.All(ctx, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (up *uploadedVideoStorage) Get(ctx context.Context, id string) (resp *models.UploadedVideoFull,
	err error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	err = up.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (up *uploadedVideoStorage) Delete(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	resp := up.collection.FindOneAndDelete(ctx, bson.M{"_id": objId})
	if resp.Err() != nil {
		return resp.Err()
	}

	return nil
}
