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

	a.log.Info("audio layers are retrieved", logger.String("path", input))

	return audioLayers, nil
}

func (a *AudioAPI) ExtractAudio(input, lang, slug, output, disk string, index int) error {
	var (
		outputPath         = fmt.Sprintf("%saudios/%s", output, lang)
		extractAudioScript = fmt.Sprintf("%s%s", a.cfg.ScriptsFolder, "/ffmpeg/extract_audio.sh")
	)

	a.log.Info("extracting audio info", logger.String("slug", slug),
		logger.String("lang", lang), logger.String("disk", disk))

	cmd, err := exec.Command(
		"/bin/sh",
		extractAudioScript,
		input,               // input path
		a.defaultBitrate,    // audio bitrate
		strconv.Itoa(index), // index of stream
		a.defaultChunkSize,  // chunk size
		outputPath,          // where to save the file
	).CombinedOutput()

	if err != nil {
		a.log.Error("failed to extract audio", logger.Error(err))
		return errors.New(string(cmd))
	}

	a.log.Info("extract output audio done")

	return nil
}
