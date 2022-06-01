package transcoder

import (
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/ffmpeg"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/audio"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/cloud"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/folder"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/subtitle"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/transcode/video"
)

type Transcoder interface {
	audio.AudioExtractor
	video.VideoExtracter
	subtitle.SubtitleExtracter
	folder.FileFolderGenerator
	cloud.ObjectUploader
}

type transcoderStruct struct {
	cfg *config.Config
	log logger.Logger

	video    video.VideoExtracter
	audio    audio.AudioExtractor
	subtitle subtitle.SubtitleExtracter
	folder   folder.FileFolderGenerator
	cloud    cloud.ObjectUploader
}

func NewTranscoder(cfg *config.Config, log logger.Logger) Transcoder {
	return &transcoderStruct{
		cfg:      cfg,
		log:      log,
		video:    video.NewVideoExtractor(cfg, log),
		audio:    audio.NewAudioExtractor(cfg, log),
		subtitle: subtitle.NewSubtitleExtracter(cfg, log),
		folder:   folder.NewFolderGenerator(cfg, log),
		cloud:    cloud.NewObjectUploader(cfg, log),
	}
}

func (t *transcoderStruct) ExtractAudio(input, lang, slug string, index int) error {
	err := t.audio.ExtractAudio(input, lang, slug, index)
	if err != nil {
		t.log.Error("failed to extract audios", logger.Error(err))
		return err
	}

	return nil
}

func (t *transcoderStruct) ExtractAudioLayers(input string) ([]ffmpeg.Stream, error) {
	streams, err := t.audio.ExtractAudioLayers(input)
	if err != nil {
		t.log.Error("failed to extract audio layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (t *transcoderStruct) ExtractInfos(input string) ([]ffmpeg.Stream, error) {
	streams, err := t.video.ExtractInfos(input)
	if err != nil {
		t.log.Info("failed to retrieve info about video layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (t *transcoderStruct) ResizeVideo(args ffmpeg.ResizeVideoArgs) error {
	err := t.video.ResizeVideo(args)
	if err != nil {
		t.log.Error("failed to resize video", logger.Error(err))
		return err
	}

	return nil
}

func (t *transcoderStruct) GetSubtitleLayers(input string) ([]ffmpeg.Stream, error) {
	streams, err := t.subtitle.GetSubtitleLayers(input)
	if err != nil {
		t.log.Error("failed to get subtitle layers", logger.Error(err))
		return streams, err
	}

	return streams, nil
}

func (t *transcoderStruct) ExtractSubtitle(input, lang, slug string, index int) error {
	err := t.subtitle.ExtractSubtitle(input, lang, slug, index)
	if err != nil {
		t.log.Error("failed to extract subtitles", logger.Error(err))
		return err
	}

	return nil
}

func (t *transcoderStruct) GenerateFilesDirectory(opts folder.FolderOpts) error {
	err := t.folder.GenerateFilesDirectory(opts)
	if err != nil {
		t.log.Error("failed to create files", logger.Error(err))
		return err
	}

	return nil
}

func (t *transcoderStruct) GenerateMasterPlaylist(opts folder.GenerateMasterOpts) error {
	err := t.folder.GenerateMasterPlaylist(opts)
	if err != nil {
		t.log.Error("failed to create master playlist")
		return err
	}

	return nil
}

func (t *transcoderStruct) UploadToS3(input, output string) error {
	err := t.cloud.UploadToS3(input, output)
	if err != nil {
		t.log.Error("failed to upload to cloud")
		return err
	}

	return nil
}
