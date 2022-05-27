package ffmpeg

type LayerOutput struct {
	Streams []Stream `json:"streams" bson:"streams"`
}

type Disposition struct {
	Default         int `json:"default" bson:"default"`
	Dub             int `json:"dub" bson:"dub"`
	Original        int `json:"original" bson:"original"`
	Comment         int `json:"comment" bson:"comment"`
	Lyrics          int `json:"lyrics" bson:"lyrics"`
	Karaoke         int `json:"karaoke" bson:"karaoke"`
	Forced          int `json:"forced" bson:"forced"`
	HearingImpaired int `json:"hearing_impaired" bson:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired" bson:"visual_impaired"`
	CleanEffects    int `json:"clean_effects" bson:"clean_effects"`
	AttachedPic     int `json:"attached_pic" bson:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails" bson:"timed_thumbnails"`
	Captions        int `json:"captions" bson:"captions"`
	Descriptions    int `json:"descriptions" bson:"descriptions"`
	Metadata        int `json:"metadata" bson:"metadata"`
	Dependent       int `json:"dependent" bson:"dependent"`
	StillImage      int `json:"still_image" bson:"still_image"`
}

type Tags struct {
	Language string `json:"language" bson:"language"`
	Title    string `json:"title" bson:"title"`
	Filename string `json:"filename" bson:"filename"`
	Mimetype string `json:"mimetype" bson:"mimetype"`
}

type Stream struct {
	Index              int         `json:"index" bson:"index"`
	CodecName          string      `json:"codec_name" bson:"codec_name"`
	CodecLongName      string      `json:"codec_long_name" bson:"codec_long_name"`
	Profile            string      `json:"profile,omitempty" bson:"profile"`
	CodecType          string      `json:"codec_type" bson:"codec_type"`
	CodecTagString     string      `json:"codec_tag_string" bson:"codec_tag_string"`
	CodecTag           string      `json:"codec_tag" bson:"codec_tag"`
	Width              int         `json:"width,omitempty" bson:"width"`
	Height             int         `json:"height,omitempty" bson:"height"`
	CodedWidth         int         `json:"coded_width,omitempty" bson:"coded_width"`
	CodedHeight        int         `json:"coded_height,omitempty" bson:"coded_height"`
	ClosedCaptions     int         `json:"closed_captions,omitempty" bson:"closed_captions"`
	FilmGrain          int         `json:"film_grain,omitempty" bson:"film_grain"`
	HasBFrames         int         `json:"has_b_frames,omitempty" bson:"has_b_frames"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio,omitempty" bson:"sample_aspect_ratio"`
	DisplayAspectRatio string      `json:"display_aspect_ratio" bson:"display_aspect_ratio"`
	PixFmt             string      `json:"pix_fmt,omitempty" bson:"pix_fmt"`
	Level              int         `json:"level,omitempty" bson:"level"`
	ChromaLocation     string      `json:"chroma_location,omitempty" bson:"chroma_location"`
	FieldOrder         string      `json:"field_order,omitempty" bson:"field_order"`
	Refs               int         `json:"refs,omitempty" bson:"refs"`
	IsAvc              string      `json:"is_avc,omitempty" bson:"is_avc"`
	NalLengthSize      string      `json:"nal_length_size,omitempty" bson:"nal_length_size"`
	RFrameRate         string      `json:"r_frame_rate" bson:"r_frame_rate"`
	AvgFrameRate       string      `json:"avg_frame_rate" bson:"avg_frame_rate"`
	TimeBase           string      `json:"time_base" bson:"time_base"`
	StartPts           int         `json:"start_pts" bson:"start_pts"`
	StartTime          string      `json:"start_time" bson:"start_time"`
	BitsPerRawSample   string      `json:"bits_per_raw_sample,omitempty" bson:"bits_per_raw_sample"`
	Disposition        Disposition `json:"disposition" bson:"disposition"`
	SampleFmt          string      `json:"sample_fmt,omitempty" bson:"sample_fmt"`
	SampleRate         string      `json:"sample_rate,omitempty" bson:"sample_rate"`
	Channels           int         `json:"channels,omitempty" bson:"channels"`
	ChannelLayout      string      `json:"channel_layout,omitempty" bson:"channel_layout"`
	BitsPerSample      int         `json:"bits_per_sample,omitempty" bson:"bits_per_sample"`
	BitRate            string      `json:"bit_rate,omitempty" bson:"bit_rate"`
	DurationTs         int         `json:"duration_ts,omitempty" bson:"duration_ts"`
	Duration           string      `json:"duration,omitempty" bson:"duration"`
	Tags               Tags        `json:"tags,omitempty" bson:"tags"`
}

type Resolution struct {
	Width   int    `json:"width,omitempty" bson:"width"`
	Height  int    `json:"height,omitempty" bson:"height"`
	BitRate string `json:"bit_rate,omitempty" bson:"bit_rate"`
}
