package interfaces

type AudioI interface {
	RetreiveAudioChannels(input string) string
	ExtractAudio(input string) error
}
