package folder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type FolderOpts struct {
	OutputPath   string
	AudioList    []ffmpeg.Tags
	SubtitleList []ffmpeg.Tags
	VideoList    string
	ScriptPath   string
}

type GenerateMasterOpts struct {
	InputPath      string
	AudioList      []ffmpeg.Tags
	SubtitleList   []ffmpeg.Tags
	ResolutionList []ffmpeg.Resolution
	Slug           string
}

type FileFolderGenerator interface {
	GenerateFilesDirectory(opts FolderOpts) error
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

func (f *FolderObject) GenerateFilesDirectory(opts FolderOpts) error {
	var (
		err               error
		scriptPath        = fmt.Sprintf("%s%s", opts.ScriptPath, "/ffmpeg/gen_folder.sh")
		audios, subtitles string
	)

	for index, audio := range opts.AudioList {
		if index == 0 {
			audios = audio.Title
		} else {
			audios = audios + "," + audio.Title
		}
	}

	for index, subtitle := range opts.SubtitleList {
		if index == 0 {
			subtitles = subtitle.Title
		} else {
			subtitles = subtitles + "," + subtitle.Title
		}
	}

	_, err = exec.Command("/bin/bash", scriptPath, opts.OutputPath, audios,
		opts.VideoList, subtitles).Output()

	return err
}

func (f *FolderObject) GenerateMasterPlaylist(opts GenerateMasterOpts) error {
	var (
		err        error
		outputPath = fmt.Sprintf("%s/%s", f.cfg.OutputDir, opts.Slug)

		fileContent     string
		stereoContent   string
		subtitleContent string
		videoContent    string
	)

	f.log.Info("Started generating master playlist")

	fileContent = masterHeader

	for index, audio := range opts.AudioList {
		stereoContent = fmt.Sprintf(stereoTemplate, stereoContent, stereoGroup, audio.Language,
			audio.Title, getLanguage(index), audio.Language)
	}

	fileContent = fileContent + "\n" + stereoContent

	for index, subtitle := range opts.SubtitleList {
		subtitleContent = fmt.Sprintf(subtitleTemplate, subtitleContent, subtitleGroup,
			subtitle.Language, subtitle.Title, getLanguage(index), subtitle.Language)
	}

	fileContent = fileContent + "\n" + subtitleContent

	for _, video := range opts.ResolutionList {
		videoContent = fmt.Sprintf(videoTemplate, videoContent, video.BitRate, codecs,
			video.Width, video.Height, stereoGroup, subtitleGroup, video.Height)
	}

	fileContent = fileContent + "\n" + videoContent

	file, err := os.Create(outputPath + "/master.m3u8")
	if err != nil {
		f.log.Error("creating master.m3u8 file error", logger.Error(err))
		return errors.New("failed at generating master playlist")
	}

	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		f.log.Error("writing to master.m3u8 file error", logger.Error(err))
		return errors.New("failed at generating master playlist")
	}

	err = file.Sync()
	if err != nil {
		f.log.Error("syncing master.m3u8 file error", logger.Error(err))
		return errors.New("failed at generating master playlist")
	}

	return nil
}

func getLanguage(index int) string {
	if index == 0 {
		return "YES"
	}

	return "NO"
}
