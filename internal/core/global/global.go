package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/core/logger"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/core/validation"
)

// global is a package that contains global variables that are used across the application.
// Init() function must be called at the application startup.

var (
	log     logger.ILogger
	logOnce sync.Once
)

func InitLog(logger logger.ILogger) {
	logOnce.Do(func() {
		log = logger
	})
}

func Log() logger.ILogger {
	return log
}

var (
	appMode     mode.Mode
	appModeOnce sync.Once
)

func InitMode(mode mode.Mode) {
	appModeOnce.Do(func() {
		appMode = mode
	})
}

func Mode() mode.Mode {
	return appMode
}

var validate validation.Validator = validation.NewValidator()

func Validator() validation.Validator {
	return validate
}
