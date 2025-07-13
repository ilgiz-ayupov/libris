package logger

import (
	"log/slog"
	"os"
)

type SlogLogger struct {
	log *slog.Logger
}

func NewSlogLogger() Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
}

func (l *SlogLogger) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args...)
}

func (l *SlogLogger) Debug(msg string, args ...interface{}) {
	l.log.Debug(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...interface{}) {
	l.log.Warn(msg, args...)
}
