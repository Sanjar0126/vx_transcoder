package main

import (
	"context"
	"net"

	"github.com/robfig/cron/v3"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/cronjob"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	transcoder "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
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

	folderGen := folder.NewFolderGenerator(&cfg, log)
	videoExr := video.NewVideoExtractor(&cfg, log)
	audioExr := audio.NewAudioExtracter(&cfg, log)
	subExr := subtitle.NewSubtitleExtracter(&cfg, log)
	//uploader := cloud.NewObjectUploader(&cfg, log)
	//var t transcoder.Transcoder

	transcodeOpts := transcoder.Opts{
		FolderGen: folderGen,
		VideoExt:  videoExr,
		AudioExt:  audioExr,
		SubExt:    subExr,
		Log:       log,
		Cfg:       &cfg,
		DB:        storageDB,
	}

	workerPool := transcoder.WorkerPools{
		JobsMap:           make(map[string]struct{}),
		MasterInfoJobs:    make(chan string, config.JobCount),
		FolderJobs:        make(chan string, config.JobCount),
		AudioJobs:         make(chan string, config.JobCount),
		VideoJobs:         make(chan string, config.JobCount),
		SubtitleJobs:      make(chan string, config.JobCount),
		ObjectStorageJobs: make(chan string, config.JobCount),
		Opts:              transcodeOpts,
	}

	go workerPool.CreateFolder()
	go workerPool.AudioInfo()
	go workerPool.SubtitleInfo()
	go workerPool.VideoInfo()

	c := cron.New()
	newCronJob := cronjob.NewCronjob(log, cfg, c, storageDB, workerPool)
	newCronJob.Run()

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
