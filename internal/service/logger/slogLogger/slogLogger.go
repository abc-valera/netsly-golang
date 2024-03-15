package slogLogger

import (
	"context"
	"log/slog"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/logger"
)

var (
	levelTrace = slog.Level(-8)
	levelFatal = slog.Level(12)

	levelNames = map[slog.Level]string{
		levelTrace: "TRACE",
		levelFatal: "FATAL",
	}
)

type slogLogger struct {
	logger *slog.Logger
}

func New() logger.ILogger {
	return &slogLogger{
		logger: slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: levelTrace,
					ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
						if a.Key == slog.LevelKey {
							level := a.Value.Any().(slog.Level)
							levelLabel, exists := levelNames[level]
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
	l.logger.Log(context.Background(), levelTrace, msg, vals)
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
	l.logger.Log(context.Background(), levelFatal, msg, vals...)
	os.Exit(1)
}
