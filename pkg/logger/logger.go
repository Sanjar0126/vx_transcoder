package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field ...
type Field = zapcore.Field

var (
	// Int ..
	Int = zap.Int
	// String ...
	String = zap.String
	// Error ...
	Error = zap.Error
	// Bool ...
	Bool = zap.Bool

	// Any ...
	Any = zap.Any
)

// Logger ...
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type loggerImpl struct {
	zap *zap.Logger
}

var (
	customTimeFormat string
)

// New ...
func New(level string, namespace string) Logger {
	if level == "" {
		level = LevelInfo
	}

	logger := loggerImpl{
		zap: newZapLogger(level, time.RFC3339),
	}

	logger.zap = logger.zap.Named(namespace)

	zap.RedirectStdLog(logger.zap)

	return &logger
}

func (log *loggerImpl) Debug(msg string, fields ...Field) {
	log.zap.Debug(msg, fields...)
}

func (log *loggerImpl) Info(msg string, fields ...Field) {
	log.zap.Info(msg, fields...)
}

func (log *loggerImpl) Warn(msg string, fields ...Field) {
	log.zap.Warn(msg, fields...)
}

func (log *loggerImpl) Error(msg string, fields ...Field) {
	log.zap.Error(msg, fields...)
}

func (log *loggerImpl) Fatal(msg string, fields ...Field) {
	log.zap.Fatal(msg, fields...)
}

// GetNamed ...
func GetNamed(log Logger, name string) Logger {
	switch v := log.(type) {
	case *loggerImpl:
		v.zap = v.zap.Named(name)
		return v
	default:
		log.Info("logger.GetNamed: invalid logger type")
		return log
	}
}

// WithFields ...
func WithFields(log Logger, fields ...Field) Logger {
	switch v := log.(type) {
	case *loggerImpl:
		return &loggerImpl{
			zap: v.zap.With(fields...),
		}
	default:
		log.Info("logger.WithFields: invalid logger type")
		return log
	}
}

// Cleanup ...
func Cleanup(log Logger) error {
	switch v := log.(type) {
	case *loggerImpl:
		return v.zap.Sync()
	default:
		log.Info("logger.Cleanup: invalid logger type")
		return nil
	}
}
