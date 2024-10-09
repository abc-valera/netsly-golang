package loggerNop

import "github.com/abc-valera/netsly-golang/internal/domain/global"

type nopLogger struct{}

// New creates a new instance of the nop logger.
// Nop means `no operation` - it's a logger that does nothing.
func New() global.ILogger {
	return &nopLogger{}
}

func (nopLogger) Trace(_ string, _ ...any) {
}

func (nopLogger) Debug(_ string, _ ...any) {
}

func (nopLogger) Info(_ string, _ ...any) {
}

func (nopLogger) Warn(_ string, _ ...any) {
}

func (nopLogger) Error(_ string, _ ...any) {
}

func (nopLogger) Fatal(_ string, _ ...any) {
}
