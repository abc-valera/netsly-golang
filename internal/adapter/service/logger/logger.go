package logger

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"golang.org/x/exp/slog"
)

var (
	LevelTrace = slog.Level(-8)
	LevelFatal = slog.Level(12)

	LevelNames = map[slog.Level]string{
		LevelTrace: "TRACE",
		LevelFatal: "FATAL",
	}
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() service.ILogger {
	return &slogLogger{
		logger: slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: LevelTrace,
					ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
						if a.Key == slog.LevelKey {
							level := a.Value.Any().(slog.Level)
							levelLabel, exists := LevelNames[level]
							if !exists {
								levelLabel = level.String()
							}

							a.Value = slog.StringValue(levelLabel)
						}

						return a
					},
				},
			),
		),
	}
}

func (l slogLogger) Trace(msg string, vals ...interface{}) {
	l.logger.Log(context.Background(), LevelTrace, msg, vals)
}

func (l slogLogger) Debug(msg string, vals ...interface{}) {
	l.logger.Debug(msg, vals...)
}

func (l slogLogger) Info(msg string, vals ...interface{}) {
	l.logger.Info(msg, vals...)
}

func (l slogLogger) Warn(msg string, vals ...interface{}) {
	l.logger.Warn(msg, vals...)
}

func (l slogLogger) Error(msg string, vals ...interface{}) {
	l.logger.Error(msg, vals...)
}

func (l slogLogger) Fatal(msg string, vals ...interface{}) {
	l.logger.Log(context.Background(), LevelFatal, msg, vals...)
	os.Exit(1)
}
