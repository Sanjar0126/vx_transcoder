package ffmpeg

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type SubtitleAPI struct {
	cfg *config.Config
	log logger.Logger

	codecType        string
	defaultChunkSize string
}

func NewSubtitleAPI(cfg *config.Config, log logger.Logger) SubtitleAPI {
	return SubtitleAPI{
		cfg:              cfg,
		log:              log,
		codecType:        "subtitle",
		defaultChunkSize: "5",
	}
}

func (s *SubtitleAPI) GetSubtitleLayers(input string) ([]Stream, error) {
	subtitleLayers, err := getLayers(s.cfg.ScriptsFolder, input, s.codecType)
	if err != nil {
		s.log.Error("failed to retrieve subtitle layers", logger.Error(err))
		return nil, err
	}

	s.log.Info("subtitle layers are retrieved")

	return subtitleLayers, nil
}

func (s *SubtitleAPI) ExtractSubtitle(input, lang, slug string, index int) error {
	var (
		outputPath            = fmt.Sprintf("%s/%s/subtitles/%s", s.cfg.OutputDir, slug, lang)
		extractSubtitleScript = fmt.Sprintf(
			"%s%s", s.cfg.ScriptsFolder, "/ffmpeg/extract_subtitle.sh")
	)

	s.log.Info("extracting audio info", logger.String("input", input), logger.String("lang", lang))

	_, err := exec.Command(
		"/bin/sh",
		extractSubtitleScript,
		input,               // input path
		strconv.Itoa(index), // index of stream
		s.defaultChunkSize,  // chunk size
		outputPath,          // where to save the output
	).Output()

	if err != nil {
		s.log.Error("failed to extract audio", logger.Error(err))
		return errors.New("failed to extract audio")
	}

	s.log.Info("extract output", logger.String("output", string("")))

	return nil
}
