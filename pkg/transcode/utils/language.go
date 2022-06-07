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
		lang = fmt.Sprintf("%s_%d", input.Language, idx+1)
	}

	if input.Title == "" {
		title = getTitle(input.Language, idx+1)
	} else {
		title = input.Title
	}

	return fffmpeg.Tags{Title: title, Language: lang}
}

func getTitle(lang string, num int) string {
	switch lang {
	case "eng":
		return "English"
	case "en":
		return "English"
	case "ru":
		return "Русский"
	case "ukr":
		return "Украинский"
	default:
		return fmt.Sprintf("Track %d", num)
	}
}
