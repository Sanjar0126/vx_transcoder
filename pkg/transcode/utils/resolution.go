package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetVideosArrayString(streams fffmpeg.Stream) string {
	var (
		resolutions = []fffmpeg.Resolution{
			// {Width: 320, Height: 240},
			// {Width: 640, Height: 360},
			{Width: 854, Height: 480},
			{Width: 1280, Height: 720},
			{Width: 1920, Height: 1080},
			{Width: 2560, Height: 1440},
			{Width: 3840, Height: 2160},
		}
		resolutionStrings []string
		ids               []int
		removed           = false
	)

	for index, resolution := range resolutions {
		if resolution.Width > streams.Width {
			ids = append(ids, index)
		}
	}

	for i, idx := range ids {
		resolutions = append(resolutions[:idx-i], resolutions[idx-i+1:]...)
		removed = true
	}

	if removed {
		resolutions = append(resolutions, fffmpeg.Resolution{
			Width:  streams.Width,
			Height: getHeight(streams.Width),
		})
	}

	for _, resolution := range resolutions {
		resolutionStrings = append(resolutionStrings, fmt.Sprintf("%dp", resolution.Height))
	}

	return strings.Join(resolutionStrings, ",")
}

func GetResolution(stream fffmpeg.Stream) []fffmpeg.Resolution {
	var (
		resolutions = []fffmpeg.Resolution{
			// {Width: 320, Height: 240},
			// {Width: 640, Height: 360},
			{Width: 854, Height: 480},
			{Width: 1280, Height: 720},
			{Width: 1920, Height: 1080},
			{Width: 2560, Height: 1440},
			{Width: 3840, Height: 2160},
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

	if removed && math.Abs(float64(stream.Width-resolutions[len(resolutions)-1].Width)) > 100 {
		resolutions = append(resolutions, fffmpeg.Resolution{
			Width:   stream.Width,
			Height:  getHeight(stream.Width),
			BitRate: GetBitRate(stream.Width, stream.Height),
		})
	}

	return resolutions
}

func getHeight(width int) int {
	if width <= 320 {
		return 240
	} else if 320 < width && width <= 640 {
		return 360
	} else if 640 < width && width <= 854 {
		return 480
	} else if 854 < width && width <= 1280 {
		return 720
	} else if 1280 < width && width <= 1920 {
		return 1080
	} else if 1920 < width && width <= 2560 {
		return 1440
	} else {
		return 2160
	}
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
		bpp = 0.078
		fps = 23.999
	)

	bitRate := float64(width) * float64(height) * bpp * fps

	return strconv.Itoa(int(bitRate))
}
