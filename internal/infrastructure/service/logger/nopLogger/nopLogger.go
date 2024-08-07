package nopLogger

import "github.com/abc-valera/netsly-api-golang/internal/domain/service"

type nopLogger struct{}

// New creates a new instance of the nop logger.
// Nop means `no operation` - it's a logger that does nothing.
func New() service.ILogger {
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
