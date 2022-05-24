package ffmpeg

import (
	"errors"
	"fmt"
	"internal/itoa"
	"os/exec"

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

	s.log.Info("subtitle layers are retreived", logger.Any("subtitles", subtitleLayers))

	return subtitleLayers, nil
}

func (s *SubtitleAPI) ExtractSubtitle(input, lang string, inputObject Stream) error {
	var (
		outputPath            = fmt.Sprintf("%s/%s/subtitles/%s", s.cfg.OutputDir, input, lang)
		extractSubtitleScript = fmt.Sprintf(
			"%s%s", s.cfg.ScriptsFolder, "/ffmpeg/extract_subtitle.sh")
	)

	s.log.Info("extracting audio info", logger.String("input", input), logger.String("lang", lang))

	out, err := exec.Command(
		"/bin/sh",
		extractSubtitleScript,
		input,                        // input path
		itoa.Itoa(inputObject.Index), // index of stream
		s.defaultChunkSize,           // chunk size
		outputPath,                   // where to save the output
	).Output()

	if err != nil {
		s.log.Error("failed to extract audio", logger.Error(err))
		return errors.New("failed to extract audio")
	}

	s.log.Info("extract output", logger.String("output", string(out)))

	return nil
}
