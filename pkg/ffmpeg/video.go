package ffmpeg

import (
	"errors"
	"fmt"
	"os/exec"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
)

type VideoAPI struct {
	cfg *config.Config
	log logger.Logger

	codecType        string
	defaultChunkSize string
}

func NewVideoAPI(cfg *config.Config, log logger.Logger) VideoAPI {
	return VideoAPI{
		cfg:              cfg,
		log:              log,
		codecType:        "video",
		defaultChunkSize: "5",
	}
}

func (v *VideoAPI) GetVideoLayers(input string) ([]Stream, error) {
	videoLayers, err := getLayers(v.cfg.ScriptsFolder, input, v.codecType)
	if err != nil {
		v.log.Error("failed to retrieve video layers", logger.Error(err))
		return nil, err
	}

	v.log.Info("video layers are retrieved", logger.Any("videos", input))

	return videoLayers, nil
}

type ResizeVideoArgs struct {
	Slug        string
	Input       string
	Width       string
	Height      string
	BitRate     string
	InputObject Stream
	OutputPath  string
	Disk        string
	Codec       string
}

func (v *VideoAPI) ResizeVideo(args ResizeVideoArgs) error {
	var (
		outputPath = fmt.Sprintf(
			"%svideos/%sp", args.OutputPath, args.Height,
		)

		resizeVideoScript   string
		resizingWidthHeight = fmt.Sprintf("%s:-2", args.Width)
	)

	if args.Codec == "hevc" || args.Codec == "h265" || args.Codec == "x265" {
		resizeVideoScript = fmt.Sprintf(
			"%s%s", v.cfg.ScriptsFolder, "/ffmpeg/265_resize.sh")
	} else {
		resizeVideoScript = fmt.Sprintf(
			"%s%s", v.cfg.ScriptsFolder, "/ffmpeg/resize_video.sh")
	}

	v.log.Info("resizing video", logger.String("slug", args.Slug), logger.String("disk", args.Disk),
		logger.String("resolution", args.Height))

	out, err := exec.Command(
		"/bin/sh",
		resizeVideoScript,
		args.Input,
		resizingWidthHeight,
		args.BitRate,
		v.defaultChunkSize,
		outputPath,
	).CombinedOutput()

	if err != nil {
		v.log.Error("failed to extract video", logger.Error(err), logger.String("std", string(out)))
		return errors.New(string(out))
	}

	v.log.Info("extract output", logger.String("output", ""))

	return nil
}
