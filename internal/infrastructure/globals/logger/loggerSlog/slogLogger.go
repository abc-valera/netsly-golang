package loggerSlog

import (
	"context"
	"log/slog"
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
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
func New(slogLoggerFolderPath string) global.ILogger {
	// Create the folder
	if err := os.MkdirAll(slogLoggerFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			coderr.Fatal(err)
		}
	}

	// Create the file inside the logs folder
	logsFile, err := os.Create(slogLoggerFolderPath + "/logs.txt")
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

func (l slogLogger) Trace(msg string, vals ...any) {
	l.stdoutLogger.Log(context.Background(), levelTrace, msg, vals...)
	l.fileLogger.Log(context.Background(), levelTrace, msg, vals...)
}

func (l slogLogger) Debug(msg string, vals ...any) {
	l.stdoutLogger.Debug(msg, vals...)
	l.fileLogger.Debug(msg, vals...)
}

func (l slogLogger) Info(msg string, vals ...any) {
	l.stdoutLogger.Info(msg, vals...)
	l.fileLogger.Info(msg, vals...)
}

func (l slogLogger) Warn(msg string, vals ...any) {
	l.stdoutLogger.Warn(msg, vals...)
	l.fileLogger.Warn(msg, vals...)
}

func (l slogLogger) Error(msg string, vals ...any) {
	l.stdoutLogger.Error(msg, vals...)
	l.fileLogger.Error(msg, vals...)
}

var handlerOptions = slog.HandlerOptions{
	Level: levelTrace,
	ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			level, ok := a.Value.Any().(slog.Level)
			if !ok {
				return a
			}

			levelLabel, exists := levelNames[level]
			if !exists {
				levelLabel = level.String()
			}

			a.Value = slog.StringValue(levelLabel)
		}

		return a
	},
}
