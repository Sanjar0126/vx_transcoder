package utils

import (
	"fmt"
	"strings"

	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetVideosArrayString(streams []fffmpeg.Stream) string {
	var (
		res         = ""
		resolutions = []int{240, 360, 480, 720, 1080}
		ids         []int
		removed     = false
	)

	for _, stream := range streams {
		for idx, resolution := range resolutions {
			if resolution > stream.Height {
				ids = append(ids, idx)
			}
		}

		for i, idx := range ids {
			resolutions = remove(resolutions, idx-i)
			removed = true
		}

		if removed {
			resolutions = append(resolutions, stream.Height)
		}
	}

	for _, resol := range resolutions {
		if res == "" {
			res = fmt.Sprintf("%dp", resol)
		} else {
			res = fmt.Sprintf("%s,%dp", res, resol)
		}
	}

	return res
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func GetLang(input string, idx int) string {
	if input == "" {
		return fmt.Sprintf("track_%d", idx)
	}

	return input
}

func GetWidth(stream fffmpeg.Stream) ([]int, []int) {
	var (
		width   = []int{320, 640, 854, 1280, 1960}
		height  = []int{240, 360, 480, 720, 1080}
		ids     []int
		removed = false
	)

	for idx, resolution := range width {
		if resolution > stream.Width {
			ids = append(ids, idx)
		}
	}

	for i, idx := range ids {
		width = remove(width, idx-i)
		height = remove(height, idx-i)
		removed = true
	}

	if removed {
		width = append(width, stream.Width)
		height = append(height, stream.Width)
	}

	return width, height
}

func GetResolution(stream fffmpeg.Stream) string {
	var (
		resolutions []string
	)

	widthList, heightList := GetWidth(stream)
	for idx, width := range widthList {
		resolutions = append(resolutions, fmt.Sprintf("%d%d", width, heightList[idx]))
	}

	return strings.Join(resolutions, ",")
}

func GetBitRate(height int) string {
	if height <= 240 {
		return "350k"
	} else if 240 < height && height <= 360 {
		return "500k"
	} else if 360 < height && height <= 480 {
		return "700k"
	} else if 480 < height && height <= 720 {
		return "1.5M"
	} else {
		return "3M"
	}
}
