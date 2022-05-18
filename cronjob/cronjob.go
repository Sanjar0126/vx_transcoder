package cronjob

import (
	"github.com/robfig/cron/v3"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	client "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/grpc"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Cronjob struct {
	log        logger.Logger
	cfg        config.Config
	cronJob    *cron.Cron
	grpcClient client.ServiceManager
	db         storage.StorageI
	transcode  transcode.TranscoderI
}

func NewCronjob(log logger.Logger, cfg config.Config,
	cron *cron.Cron, db storage.StorageI) *Cronjob {
	grpcClient, err := client.NewGrpcClients(&cfg)
	if err != nil {
		panic(err)
	}

	return &Cronjob{
		cfg:        cfg,
		log:        log,
		cronJob:    cron,
		grpcClient: grpcClient,
		db:         db,
		transcode:  transcode.New(&cfg, log, db),
	}
}

func (c *Cronjob) Run() {
	_, err := c.cronJob.AddFunc("@every 1m", c.transcode.Run)
	if err != nil {
		c.log.Error("failed to register cronjob", logger.Error(err))
		panic(err)
	}

	c.log.Info("cronjob is registered")

	c.cronJob.Start()
}
