package folder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type FileFolderGenerator interface {
	GenerateFilesDirecetory()
	GenerateMasterPlaylist()
}

type FolderObject struct {
	cfg *config.Config
	log logger.Logger
}

func NewFolderGenerator(cfg *config.Config, log logger.Logger) FileFolderGenerator {
	return &FolderObject{
		cfg: cfg,
		log: log,
	}
}

func (f *FolderObject) GenerateFilesDirecetory() {

}

func (f *FolderObject) GenerateMasterPlaylist() {

}
