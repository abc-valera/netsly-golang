package logger

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"golang.org/x/exp/slog"
)

type slogLogger struct {
	log *slog.Logger
}

func NewSlogLogger() service.Logger {
	return &slogLogger{
		log: slog.Default(),
	}
}

func (l *slogLogger) Debug(msg string, vals ...interface{}) {
	l.log.Debug(msg, vals...)
}

func (l *slogLogger) Info(msg string, vals ...interface{}) {
	l.log.Info(msg, vals...)
}

func (l *slogLogger) Warn(msg string, vals ...interface{}) {
	l.log.Warn(msg, vals...)
}

func (l *slogLogger) Error(msg string, vals ...interface{}) {
	l.log.Error(msg, vals...)
}
