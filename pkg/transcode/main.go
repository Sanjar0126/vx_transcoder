package transcoder

type WorkerPools struct {
	jobMap            map[string]struct{}
	InfoJobs          chan string
	FolderJobs        chan string
	AudioJobs         chan string
	VideoJobs         chan string
	SubtitleJobs      chan string
	ObjectStorageJobs chan string
}
