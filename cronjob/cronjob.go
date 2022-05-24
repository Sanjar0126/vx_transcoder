package cronjob

import (
	"github.com/robfig/cron/v3"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type Cronjob struct {
	log     logger.Logger
	cfg     config.Config
	cronJob *cron.Cron
	db      storage.StorageI
}

func NewCronjob(log logger.Logger, cfg config.Config,
	cron *cron.Cron, db storage.StorageI) *Cronjob {
	return &Cronjob{
		cfg:     cfg,
		log:     log,
		cronJob: cron,
		db:      db,
	}
}

func (c *Cronjob) Run() {
	c.log.Info("cronjob is registered")
	c.cronJob.Start()
}
