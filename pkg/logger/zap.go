package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(level, timeFormat string) *zap.Logger {
	globalLevel := parseLevel(level)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= globalLevel && lvl < zapcore.ErrorLevel
	})

	consoleInfos := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Configure console output.
	encoderCfg := zap.NewProductionEncoderConfig()

	if len(timeFormat) > 0 {
		customTimeFormat = timeFormat
		encoderCfg.EncodeTime = customTimeEncoder
	} else {
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
	)

	logger := zap.New(core)

	return logger
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// GetZapLogger extracts zap struct from given logger interface
func GetZapLogger(log Logger) *zap.Logger {
	if log == nil {
		return newZapLogger(LevelInfo, time.RFC3339)
	}

	switch v := log.(type) {
	case *loggerImpl:
		return v.zap
	default:
		log.Info(
			"logger.WithFields: invalid logger type, creating a new zap logger",
			String("level", LevelInfo), String("time_format", time.RFC3339),
		)

		return newZapLogger(LevelInfo, time.RFC3339)
	}
}
