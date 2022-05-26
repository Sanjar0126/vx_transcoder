package ffmpeg

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type AudioAPI struct {
	cfg *config.Config
	log logger.Logger

	codecType        string
	defaultBitrate   string
	defaultChunkSize string
}

func NewAudioAPI(cfg *config.Config, log logger.Logger) AudioAPI {
	return AudioAPI{
		cfg:              cfg,
		log:              log,
		codecType:        "audio",
		defaultBitrate:   "128k",
		defaultChunkSize: "5",
	}
}

func (a *AudioAPI) GetAudioLayers(input string) ([]Stream, error) {
	audioLayers, err := getLayers(a.cfg.ScriptsFolder, input, a.codecType)
	if err != nil {
		a.log.Error("failed to retrieve audio layers", logger.Error(err))
		return nil, err
	}

	a.log.Info("audio layers are retrieved")

	return audioLayers, nil
}

func (a *AudioAPI) ExtractAudio(input, lang, slug string, index int) error {
	var (
		outputPath         = fmt.Sprintf("%s/%s/audios/%s", a.cfg.OutputDir, slug, lang)
		extractAudioScript = fmt.Sprintf("%s%s", a.cfg.ScriptsFolder, "/ffmpeg/extract_audio.sh")
	)

	a.log.Info("extracting audio info", logger.String("input", input), logger.String("lang", lang))

	cmd, err := exec.Command(
		"/bin/sh",
		extractAudioScript,
		input,               // input path
		a.defaultBitrate,    // audio bitrate
		strconv.Itoa(index), // index of stream
		a.defaultChunkSize,  // chunk size
		outputPath,          // where to save the file
	).Output()

	if err != nil {
		a.log.Error("failed to extract audio", logger.Error(err))
		return errors.New("failed to extract audio")
	}

	a.log.Info("extract output", logger.String("output", string(cmd)))

	return nil
}
