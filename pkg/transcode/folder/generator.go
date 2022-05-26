package folder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type FileFolderGenerator interface {
	GenerateFilesDirectory(input string, inputObjects []ffmpeg.Stream) error
	GenerateMasterPlaylist(input string, inputObjects []ffmpeg.Stream) error
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

func (f *FolderObject) GenerateFilesDirectory(input string, inputObjects []ffmpeg.Stream) error {
	return nil
}

func (f *FolderObject) GenerateMasterPlaylist(input string, inputObjects []ffmpeg.Stream) error {
	return nil
}
