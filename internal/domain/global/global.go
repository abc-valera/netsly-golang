package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/domain/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

// global is a package that contains global variables that are used across the application.
// Init() function must be called at the application startup.

func Mode() mode.Mode {
	return appMode
}

var appMode mode.Mode

func Log() service.ILogger {
	return log
}

var log service.ILogger

func Init(
	mode mode.Mode,
	logger service.ILogger,
) {
	initOnce.Do(func() {
		appMode = mode
		log = logger
	})
}

var initOnce sync.Once
