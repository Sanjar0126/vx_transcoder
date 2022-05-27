package utils

import (
	"fmt"

	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetAudioList(streams []fffmpeg.Stream) []fffmpeg.Tags {
	var (
		result []fffmpeg.Tags
	)

	for idx, stream := range streams {
		result = append(result, GetTag(stream.Tags, idx))
	}

	return result
}

func GetLangArrayString(streams []fffmpeg.Stream) []fffmpeg.Tags {
	var (
		lang []fffmpeg.Tags
	)

	for idx, stream := range streams {
		lang = append(lang, GetTag(stream.Tags, idx))
	}

	return lang
}

func GetTag(input fffmpeg.Tags, idx int) fffmpeg.Tags {
	var (
		title, lang string
	)

	if input.Language == "" {
		lang = fmt.Sprintf("track_%d", idx+1)
	} else {
		lang = input.Language
	}

	if input.Title == "" {
		title = fmt.Sprintf("Track %d", idx+1)
	} else {
		title = input.Title
	}

	return fffmpeg.Tags{Title: title, Language: lang}
}
