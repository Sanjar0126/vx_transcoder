package transcode

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/models"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/interfaces"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/storage"
)

type transcoder struct {
	audio    interfaces.AudioI
	video    interfaces.VideoI
	subtitle interfaces.SubtitleI
	playlist interfaces.PlaylistI
	disk     interfaces.DiskI
	db       storage.StorageI
	log      logger.Logger
}

func (t *transcoder) Run() {
	newUploadedVideos, err := t.db.UploadedVideo().GetAll(context.Background(),
		models.UploadedVideoFilter{
			Stage: "new",
		})

	if err != nil {
		t.log.Error("Error while getting new uploaded video")
		return
	}

	extractAudiosItems, err := t.db.UploadedVideo().GetAll(context.Background(),
		models.UploadedVideoFilter{
			Stage: "extract_audio",
		})
	if err != nil {
		t.log.Error("Error while getting new uploaded video")
		return
	}

	basic := make(chan string, len(newUploadedVideos))
	resultsNewUploaded := make(chan string, len(newUploadedVideos))

	extractAudioChan := make(chan string, len(extractAudiosItems))
	resultsAudioChan := make(chan string, len(extractAudiosItems))

	for k := 1; k <= 3; k++ {
		go t.WorkerNewUploaded(k, basic, resultsNewUploaded)
		go t.WorkerExtractAudio(k, extractAudioChan, resultsAudioChan)
	}

	for i := 0; i < len(newUploadedVideos); i++ {
		basic <- newUploadedVideos[i].ID
	}

	for z := 0; z < len(extractAudiosItems); z++ {
		extractAudioChan <- extractAudiosItems[z].ID
	}

	for z := 0; z < len(newUploadedVideos); z++ {
		<-resultsNewUploaded
	}

	for n := 0; n < len(extractAudiosItems); n++ {
		<-resultsAudioChan
	}

}

func (t *transcoder) WorkerNewUploaded(id int, job <-chan string, results chan<- string) {
	for j := range job {
		fmt.Println("Worker new upload", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker new upload", id, "finished job", j)

		err := t.db.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
			ID:    j,
			Stage: "extract_audio",
		})
		if err != nil {
			panic(err)
		}

		results <- j
	}
}

func (t *transcoder) WorkerExtractAudio(id int, job <-chan string, results chan<- string) {
	for j := range job {
		fmt.Println("Worker extract audio", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker extract audio", id, "finished job", j)

		err := t.db.UploadedVideo().Update(context.Background(), models.UploadVideoRequest{
			ID:    j,
			Stage: "extract_subtitle",
		})

		if err != nil {
			panic(err)
		}

		results <- j
	}
}
