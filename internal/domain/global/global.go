package global

import (
	"sync"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/core/validator"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

// global is a package that contains global variables that are used across the application.

var (
	appMode         mode.Mode
	appModeInitOnce sync.Once
)

func InitMode(m mode.Mode) {
	if m != mode.Development && m != mode.Production {
		coderr.Fatal("Provided invalid application mode. Should be either 'development' or 'production'")
	}
	appModeInitOnce.Do(func() {
		appMode = m
	})
}

func Mode() mode.Mode {
	return appMode
}

var (
	log         service.ILogger
	logInitOnce sync.Once
)

func InitLog(l service.ILogger) {
	logInitOnce.Do(func() {
		log = l
	})
}

func Log() service.ILogger {
	return log
}

var validate validator.IValidator = validator.NewValidator()

func Validate() validator.IValidator {
	return validate
}
