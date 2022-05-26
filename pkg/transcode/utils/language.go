package utils

import (
	fffmpeg "gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

func GetLangArrayString(streams []fffmpeg.Stream) string {
	var (
		res  = ""
		lang string
	)

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
