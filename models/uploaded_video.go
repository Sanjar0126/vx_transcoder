package models

import (
	"time"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
)

const (
	IdLiteral = "_id"
)

type UploadVideoRequest struct {
	ID            string `json:"id" bson:"id,omitempty"`
	MovieSlug     string `json:"movie_slug" bson:"movie_slug,omitempty"`
	Type          string `json:"type" bson:"type,omitempty"`
	SeasonNumber  int32  `json:"season_number" bson:"season_number,omitempty"`
	EpisodeNumber int32  `json:"episode_number" bson:"episode_number,omitempty"`
	Stage         string `json:"stage" bson:"stage,omitempty"`
	Path          string `json:"path" bson:"path,omitempty"`
	Failed        bool   `bson:"failed,omitempty"`
	Extension     string `json:"extension" bson:"extension"`
}

type UpdateStreams struct {
	AudioStreams    []ffmpeg.Stream `json:"audio_streams,omitempty" bson:"audio_streams"`
	VideoStreams    []ffmpeg.Stream `json:"video_streams,omitempty" bson:"video_streams"`
	SubtitleStreams []ffmpeg.Stream `json:"subtitle_streams,omitempty" bson:"subtitle_streams"`
	Stage           string          `json:"stage,omitempty" bson:"stage"`
	ID              string          `json:"id,omitempty" bson:"id"`
}

type UploadedVideo struct {
	FilePath string `json:"file_path" bson:"file_path"`
}

type UploadedVideoFilter struct {
	Stages []string `json:"stages" bson:"stages,omitempty"`
}

type UploadedVideoFull struct {
	ID            string          `json:"_id" bson:"_id,omitempty"` //nolint
	MovieSlug     string          `json:"movie_slug" bson:"movie_slug,omitempty"`
	Type          string          `json:"type" bson:"type,omitempty"`
	SeasonNumber  int32           `json:"season_number" bson:"season_number,omitempty"`
	EpisodeNumber int32           `json:"episode_number" bson:"episode_number,omitempty"`
	Stage         string          `json:"stage" bson:"stage,omitempty"`
	Path          string          `json:"path" bson:"path,omitempty"`
	Extension     string          `json:"extension" bson:"extension"`
	Audios        []ffmpeg.Stream `json:"audios,omitempty" bson:"audios"`
	Videos        []ffmpeg.Stream `json:"videos,omitempty" bson:"videos"`
	Subtitles     []ffmpeg.Stream `json:"subtitles,omitempty" bson:"subtitles"`
	UpdatedAt     time.Time       `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt     time.Time       `bson:"created_at,omitempty" json:"created_at"`
	Failed        bool            `bson:"failed"`
}
