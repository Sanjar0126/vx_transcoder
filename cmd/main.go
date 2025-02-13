package main

import (
	"context"
	"net"

	"github.com/robfig/cron/v3"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/cronjob"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/worker"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage/db"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "transcoder_service")

	log.Info("configuration and logger is setup....")

	dbConn := db.NewConn(&cfg, log)
	defer func() {
		err := dbConn.MongoConn.Client().Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	storageDB := storage.New(dbConn.MongoConn)

	transcoderObj := transcoder.NewTranscoder(&cfg, log)

	workerPool := make(map[string]worker.Worker)

	for _, disk := range config.DiskArray {
		workerPool[disk] = worker.NewWorker(transcoderObj, log, &cfg, storageDB)

		for i := 0; i < config.JobCount; i++ {
			go workerPool[disk].CreateFolder()
			go workerPool[disk].AudioInfo()
			go workerPool[disk].SubtitleInfo()
			go workerPool[disk].MasterGenerate()
		}

		for i := 0; i < config.VideoResizeJobCount; i++ {
			go workerPool[disk].VideoInfo()
		}

		for i := 0; i < config.UploadJobCount; i++ {
			go workerPool[disk].Upload()
		}
	}

	c := cron.New()
	newCronJob := cronjob.NewCronjob(log, cfg, c, storageDB, workerPool)
	newCronJob.Run()

	newCronJob.Initial()

	log.Info("main: server is running", logger.String("port", cfg.RPCPort))

	server := grpc.NewServer()

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("Error while listening: %v", logger.Error(err))
		panic(err)
	}

	if err := server.Serve(listen); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		panic(err)
	}
}
