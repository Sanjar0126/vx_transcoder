package folder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
)

func TestMasterTest(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "transcoder_service")

	fileGen := folder.NewFolderGenerator(&cfg, log)

	err := fileGen.GenerateMasterPlaylist(folder.GenerateMasterOpts{
		InputPath: "",
		AudioList: []ffmpeg.Tags{{
			Language: "ru",
			Title:    "ru-title",
		}, {
			Language: "en",
			Title:    "en-title",
		}},
		SubtitleList: []ffmpeg.Tags{{
			Language: "ru",
			Title:    "ru-title",
		}, {
			Language: "en",
			Title:    "en-title",
		}},
		ResolutionList: []ffmpeg.Resolution{{
			Width:   1920, //nolint
			Height:  1080, //nolint
			BitRate: "3000000",
		}, {
			Width:   1000, //nolint
			Height:  500,  //nolint
			BitRate: "102012",
		}},
		Slug: "bear",
	})

	assert.NoError(t, err)
}
