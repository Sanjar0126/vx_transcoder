package folder

import (
	"errors"
	"fmt"
	"os/exec"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type GenerateMasterOpts struct {
	InputPath      string
	AudioList      string
	SubtitleList   string
	ResolutionList string
	Slug           string
	Qualities      string
}

type FileFolderGenerator interface {
	GenerateFilesDirectory(input string, inputObjects []ffmpeg.Stream) error
	GenerateMasterPlaylist(opts GenerateMasterOpts) error
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

func (f *FolderObject) GenerateMasterPlaylist(opts GenerateMasterOpts) error {
	f.log.Info("Started generating master playlist")

	outputPath := fmt.Sprintf("%s/%s", f.cfg.OutputDir, opts.Slug)
	script := fmt.Sprintf("%s%s", f.cfg.ScriptsFolder, "/ffmpeg/generate_master_playlist.sh")

	_, err := exec.Command(
		"/bin/bash",
		script,
		outputPath,
		opts.AudioList,
		opts.SubtitleList,
		opts.Qualities,
		opts.ResolutionList,
	).Output()

	if err != nil {
		f.log.Error("generating master playlist err", logger.Error(err))
		return errors.New("failed at generating master playlist err")
	}

	return nil
}
