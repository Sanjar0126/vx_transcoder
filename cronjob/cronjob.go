package cronjob

import (
	"context"

	"github.com/robfig/cron/v3"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/models"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/worker"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Cronjob struct {
	log     logger.Logger
	cfg     config.Config
	cronJob *cron.Cron
	db      storage.StorageI
	worker  map[string]worker.Worker
}

func NewCronjob(log logger.Logger, cfg config.Config,
	cron *cron.Cron, db storage.StorageI, worker map[string]worker.Worker) *Cronjob {
	return &Cronjob{
		cfg:     cfg,
		log:     log,
		cronJob: cron,
		db:      db,
		worker:  worker,
	}
}

func (c *Cronjob) Run() {
	_, err := c.cronJob.AddFunc("@every 5m", c.transcode)
	if err != nil {
		c.log.Error("failed to register cronjob", logger.Error(err))
		panic(err)
	}

	_, err = c.cronJob.AddFunc("@every 5m", c.upload)
	if err != nil {
		c.log.Error("failed to register cronjob", logger.Error(err))
		panic(err)
	}

	c.log.Info("cronjob is registered")
	c.cronJob.Start()
}

func (c *Cronjob) Initial() {
	c.log.Info("cronjob is registered")
	c.transcode()
	c.upload()
}

func (c *Cronjob) upload() {
	for _, disk := range config.DiskArray {
		dbRes, err := c.db.UploadedVideo().GetAll(context.Background(), models.UploadedVideoFilter{
			Stages: []string{config.UploadStage},
			Disk:   disk,
		})
		if err != nil {
			c.log.Error("error while getting list of videos", logger.Error(err))
		}

		go c.worker[disk].DistributeJobs(dbRes)
	}
}

func (c *Cronjob) transcode() {
	for _, disk := range config.DiskArray {
		// c.log.Info("transcode cronjob", logger.String("disk", disk))

		dbRes, err := c.db.UploadedVideo().GetAll(context.Background(), models.UploadedVideoFilter{
			Stages: config.StagesArray,
			Disk:   disk,
		})
		if err != nil {
			c.log.Error("error while getting list of videos", logger.Error(err))
		}

		go c.worker[disk].DistributeJobs(dbRes)
	}
}
