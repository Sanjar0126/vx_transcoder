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

	inserted, err := up.collection.InsertOne(ctx, models.UploadedVideoFull{
		ID:            req.ID,
		MovieSlug:     req.MovieSlug,
		Type:          req.Type,
		Stage:         req.Stage,
		SeasonNumber:  req.SeasonNumber,
		EpisodeNumber: req.EpisodeNumber,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Failed:        false,
	})

	if err != nil {
		return resp.FilePath, err
	}

	resp.FilePath = inserted.InsertedID.(string)

	return resp.FilePath, nil
}

func (up *uploadedVideoStorage) Update(ctx context.Context, req models.UploadVideoRequest) error {
	var (
		update = bson.M{}
	)

	objId, objErr := primitive.ObjectIDFromHex(req.ID)
	if objErr != nil {
		return objErr
	}

	if req.Stage != "" {
		update = bson.M{"$set": bson.M{"stage": req.Stage}}
	}

	if req.Failed {
		update = bson.M{"$set": bson.M{"failed": true}}
	}

	err := up.collection.FindOneAndUpdate(ctx, bson.M{models.IdLiteral: objId}, update)

	return err.Err()
}

func (up *uploadedVideoStorage) UpdateStreams(ctx context.Context, req models.UpdateStreams) error {
	objId, objErr := primitive.ObjectIDFromHex(req.ID)
	if objErr != nil {
		return objErr
	}

	err := up.collection.FindOneAndUpdate(ctx, bson.M{models.IdLiteral: objId},
		bson.M{"$set": bson.M{
			"stage":     req.Stage,
			"audios":    req.AudioStreams,
			"subtitles": req.SubtitleStreams,
			"videos":    req.VideoStreams,
		}})

	return err.Err()
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
		"failed": false,
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
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return resp, err
	}

	err = up.collection.FindOne(ctx, bson.M{models.IdLiteral: objId}).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (up *uploadedVideoStorage) Delete(ctx context.Context, id string) error {
	objId, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		return objErr
	}

	resp := up.collection.FindOneAndDelete(ctx, bson.M{models.IdLiteral: objId})
	if resp.Err() != nil {
		return resp.Err()
	}

	return nil
}
