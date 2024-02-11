package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/domain/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

// global is a package that contains global variables that are used across the application.
// Init() function must be called at the application startup.

var (
	log     service.ILogger
	logOnce sync.Once
)

func InitLog(logger service.ILogger) {
	logOnce.Do(func() {
		log = logger
	})
}

func Log() service.ILogger {
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

var (
	validate     service.IValidator
	validateOnce sync.Once
)

func InitValidator(validator service.IValidator) {
	validateOnce.Do(func() {
		validate = validator
	})
}

func Validator() service.IValidator {
	return validate
}
