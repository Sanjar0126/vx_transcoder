package logger

var (
	levelDebug = -1
	levelInfo  = 0
	levelWarn  = 1
	levelError = 2
	levelPanic = 3
	levelFatal = 5
)

// LogLevelFromString ...
func LogLevelFromString(level string) int {
	switch level {
	case LevelDebug:
		return levelDebug
	case LevelInfo:
		return levelInfo
	case LevelWarn:
		return levelWarn
	case LevelError:
		return levelError
	case LevelPanic:
		return levelPanic
	case LevelFatal:
		return levelFatal
	default:
		return 0
	}
}
