package global

import "github.com/abc-valera/flugo-api-golang/internal/core/domain/service"

// Mode is the application running mode.
// It can be either "development" or "production".
//
// It should be set at the startup of the application.
var Mode string

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

// Log is the global the application logger.
//
// It should be set at the startup of the application.
var Log service.ILogger
