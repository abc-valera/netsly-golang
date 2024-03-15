package nopLogger

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/logger"
)

type nopLogger struct{}

func New() logger.ILogger {
	return &nopLogger{}
}

func (l nopLogger) Trace(msg string, vals ...interface{}) {
}

func (l nopLogger) Debug(msg string, vals ...interface{}) {
}

func (l nopLogger) Info(msg string, vals ...interface{}) {
}

func (l nopLogger) Warn(msg string, vals ...interface{}) {
}

func (l nopLogger) Error(msg string, vals ...interface{}) {
}

func (l nopLogger) Fatal(msg string, vals ...interface{}) {
}
