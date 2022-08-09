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

	s.log.Info("subtitle layers are retrieved", logger.String("path", input))

	return subtitleLayers, nil
}

func (s *SubtitleAPI) ExtractSubtitle(input, lang, slug, output string, index int) error {
	var (
		outputPath            = fmt.Sprintf("%ssubtitles/%s", output, lang)
		extractSubtitleScript = fmt.Sprintf(
			"%s%s", s.cfg.ScriptsFolder, "/ffmpeg/extract_subtitle.sh")
	)

	s.log.Info("extracting subtitle info", logger.String("slug", slug), logger.String("lang", lang))

	out, err := exec.Command(
		"/bin/sh",
		extractSubtitleScript,
		input,               // input path
		strconv.Itoa(index), // index of stream
		s.defaultChunkSize,  // chunk size
		outputPath,          // where to save the output
	).Output()

	if err != nil {
		s.log.Error("failed to extract audio", logger.Error(err))
		return errors.New(string(out))
	}

	s.log.Info("extract output", logger.String("output", string(out)))

	return nil
}
