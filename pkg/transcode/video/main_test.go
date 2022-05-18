package video_test

import (
	"context"
	"os"
	"testing"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage/db"
)

var (
	videoTest interfaces.VideoI
	log       logger.Logger
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	log = logger.New("debug", "transcoder-service")

	dbConn := db.NewConn(&cfg, log)
	defer func() {
		err := dbConn.MongoConn.Client().Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	storageDB := storage.New(dbConn.MongoConn)

	videoTest = video.NewVideo(&cfg, log, storageDB)

	os.Exit(m.Run())
}
