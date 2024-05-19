package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/core/validator"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
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
		validate = validator.NewValidator()
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

var validate validator.IValidator

func Validate() validator.IValidator {
	return validate
}
