package slogLogger

import (
	"context"
	"log/slog"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

var (
	levelTrace = slog.Level(-8)

	levelNames = map[slog.Level]string{
		levelTrace: "TRACE",
	}
)

type slogLogger struct {
	stdoutLogger *slog.Logger
	fileLogger   *slog.Logger
}

// New returns a new instance of the slogLogger.
// SlogLogger is a simple wrapper for a slog library that writes to stdout and to a file.
func New(logsFolderPath string) service.ILogger {
	// Create the file inside the logs folder
	logsFile, err := os.Create(logsFolderPath + "/logs.txt")
	if err != nil {
		coderr.Fatal(err)
	}

	return &slogLogger{
		stdoutLogger: slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&handlerOptions,
			),
		),
		fileLogger: slog.New(
			slog.NewJSONHandler(
				logsFile,
				&handlerOptions,
			),
		),
	}
}

func (l slogLogger) Trace(msg string, vals ...interface{}) {
	l.stdoutLogger.Log(context.Background(), levelTrace, msg, vals...)
	l.fileLogger.Log(context.Background(), levelTrace, msg, vals...)
}

func (l slogLogger) Debug(msg string, vals ...interface{}) {
	l.stdoutLogger.Debug(msg, vals...)
	l.fileLogger.Debug(msg, vals...)
}

func (l slogLogger) Info(msg string, vals ...interface{}) {
	l.stdoutLogger.Info(msg, vals...)
	l.fileLogger.Info(msg, vals...)
}

func (l slogLogger) Warn(msg string, vals ...interface{}) {
	l.stdoutLogger.Warn(msg, vals...)
	l.fileLogger.Warn(msg, vals...)
}

func (l slogLogger) Error(msg string, vals ...interface{}) {
	l.stdoutLogger.Error(msg, vals...)
	l.fileLogger.Error(msg, vals...)
}

var handlerOptions = slog.HandlerOptions{
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
}
