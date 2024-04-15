package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/pkg/core/mode"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
)

// global is a package that contains global variables that are used across the application.
// Init() function must be called at the application startup.

var initOnce sync.Once

func Init(
	mode mode.Mode,
	logger service.ILogger,
) {
	initOnce.Do(func() {
		appMode = mode
		log = logger
	})
}

var appMode mode.Mode

func Mode() mode.Mode {
	return appMode
}

var log service.ILogger

func Log() service.ILogger {
	return log
}
