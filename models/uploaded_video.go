package models

import (
	"time"
)

type UploadVideoRequest struct {
	ID            string `json:"id" bson:"id,omitempty"`
	MovieSlug     string `json:"movie_slug" bson:"movie_slug,omitempty"`
	Type          string `json:"type" bson:"type,omitempty"`
	SeasonNumber  int32  `json:"season_number" bson:"season_number,omitempty"`
	EpisodeNumber int32  `json:"episode_number" bson:"episode_number,omitempty"`
	Stage         string `json:"stage" bson:"stage,omitempty"`
}

type UploadedVideo struct {
	FilePath string `json:"file_path" bson:"file_path"`
}

type UploadedVideoFilter struct {
	Stage string `json:"stage" bson:"stage,omitempty"`
}

type UploadedVideoFull struct {
	ID            string    `json:"id" bson:"id,omitempty"`
	MovieSlug     string    `json:"movie_slug" bson:"movie_slug,omitempty"`
	Type          string    `json:"type" bson:"type,omitempty"`
	SeasonNumber  int32     `json:"season_number" bson:"season_number,omitempty"`
	EpisodeNumber int32     `json:"episode_number" bson:"episode_number,omitempty"`
	Stage         string    `json:"stage" bson:"stage"`
	UpdatedAt     time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt     time.Time `bson:"created_at,omitempty" json:"created_at"`
}
