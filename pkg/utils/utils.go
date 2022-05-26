package utils

import (
	"fmt"

	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetLangArrayString(streams []fffmpeg.Stream) string {
	var res = ""
	var lang string

	for idx, stream := range streams {
		lang = GetLang(stream.Tags.Language, idx)

		if res == "" {
			res = lang
		} else {
			res = res + "," + lang
		}
	}

	return res
}

func GetVideosArrayString(streams []fffmpeg.Stream) string {
	var res = ""
	var resolutions = []int{240, 360, 480, 720, 1080}
	var ids []int
	var removed = false

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
	} else {
		return input
	}
}

func GetWidth(stream fffmpeg.Stream) []int {
	var width = []int{320, 640, 854, 1280, 1960}

	var ids []int
	var removed = false

	for idx, resolution := range width {
		if resolution > stream.Width {
			ids = append(ids, idx)
		}
	}

	for i, idx := range ids {
		width = remove(width, idx-i)
		removed = true
	}

	if removed {
		width = append(width, stream.Width)
	}

	return width
}

func GetBitRate(width int) string {
	if width <= 240 {
		return "350k"
	} else if 240 < width && width <= 360 {
		return "500k"
	} else if 360 < width && width <= 480 {
		return "700k"
	} else if 480 < width && width <= 720 {
		return "1.5M"
	} else {
		return "3M"
	}
}
