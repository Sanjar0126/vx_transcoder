package interfaces

type SubtitleI interface {
	RetreiveSubtitleChannels(input string) string
	ExtractSubtitle(input string) error
}
