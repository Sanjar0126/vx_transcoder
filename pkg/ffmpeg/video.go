package ffmpeg

import (
	"bytes"
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

	v.log.Info("video layers are retrieved", logger.Any("videos", videoLayers))

	return videoLayers, nil
}

type ResizeVideoArgs struct {
	Slug        string
	Input       string
	Width       string
	Height      string
	BitRate     string
	InputObject Stream
}

func (v *VideoAPI) ResizeVideo(args ResizeVideoArgs) error {
	var (
		outputPath = fmt.Sprintf(
			"%s/%s/videos/%sp/video.m3u8", v.cfg.OutputDir, args.Slug, args.Height,
		)

		resizeVideoScript   = fmt.Sprintf("%s%s", v.cfg.ScriptsFolder, "/ffmpeg/resize_video.sh")
		resizingWidthHeight = fmt.Sprintf("%s:-2", args.Width)
	)

	fmt.Println(outputPath)

	cmd := exec.Command(
		"/bin/sh",
		resizeVideoScript,
		args.Input,
		resizingWidthHeight,
		args.BitRate,
		v.defaultChunkSize,
		outputPath,
	)

	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	fmt.Println(string(stderr.Bytes()))

	if err != nil {
		v.log.Error("failed to extract video", logger.Error(err))
		return errors.New("failed to extract video")
	}

	v.log.Info("extract output", logger.String("output", ""))

	return nil
}
