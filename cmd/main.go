package main

import (
	"context"
	"net"

	"github.com/robfig/cron/v3"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/cronjob"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
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

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("Error while listening: %v", logger.Error(err))
		panic(err)
	}

	c := cron.New()
	newCronJob := cronjob.NewCronjob(log, cfg, c, storageDB)
	newCronJob.Run()

	log.Info("main: server is running", logger.String("port", cfg.RPCPort))

	server := grpc.NewServer()

	if err := server.Serve(listen); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		panic(err)
	}
}
