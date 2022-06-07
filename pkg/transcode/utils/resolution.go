package utils

import (
	"fmt"
	"strconv"
	"strings"

	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetVideosArrayString(streams []fffmpeg.Stream) string {
	var (
		resolutions = []fffmpeg.Resolution{
			{Width: 320, Height: 240},
			{Width: 640, Height: 360},
			{Width: 854, Height: 480},
			{Width: 1280, Height: 720},
			{Width: 1920, Height: 1080},
		}
		resolutionStrings []string
		ids               []int
		removed           = false
	)

	for _, stream := range streams {
		for index, resolution := range resolutions {
			if resolution.Width > stream.Width {
				ids = append(ids, index)
			}
		}

		for i, idx := range ids {
			resolutions = append(resolutions[:idx-i], resolutions[idx-i+1:]...)
			removed = true
		}

		if removed {
			resolutions = append(resolutions, fffmpeg.Resolution{
				Width:  stream.Width,
				Height: stream.Height,
			})
		}
	}

	for _, resolution := range resolutions {
		resolutionStrings = append(resolutionStrings, fmt.Sprintf("%dp", resolution.Height))
	}

	return strings.Join(resolutionStrings, ",")
}

func GetResolution(stream fffmpeg.Stream) []fffmpeg.Resolution {
	var (
		resolutions = []fffmpeg.Resolution{
			{Width: 320, Height: 240},
			{Width: 640, Height: 360},
			{Width: 854, Height: 480},
			{Width: 1280, Height: 720},
			{Width: 1920, Height: 1080},
		}
		ids     []int
		removed = false
	)

	for idx, resolution := range resolutions {
		if resolution.Width > stream.Width {
			ids = append(ids, idx)
		} else {
			resolutions[idx].BitRate = GetBitRate(resolution.Width, resolution.Height)
		}
	}

	for i, id := range ids {
		resolutions = append(resolutions[:id-i], resolutions[id-i+1:]...)
		removed = true
	}

	if removed {
		resolutions = append(resolutions, fffmpeg.Resolution{
			Width:   stream.Width,
			Height:  stream.Height,
			BitRate: GetBitRate(stream.Width, stream.Height),
		})
	}

	return resolutions
}

func GetBitRate(width, height int) string {
	//if height <= 240 {
	//	return "358400"
	//} else if 240 < height && height <= 360 {
	//	return "512000"
	//} else if 360 < height && height <= 480 {
	//	return "716800"
	//} else if 480 < height && height <= 720 {
	//	return "1572864"
	//} else {
	//	return "3145728"
	//}
	var (
		bpp = 0.067
		fps = 23.9
	)

	bitRate := float64(width) * float64(height) * bpp * fps

	return strconv.Itoa(int(bitRate))
}
