package app

import "errors"

// Mode is the application running mode
type Mode int

const (
	ModeDevelopment Mode = iota + 1
	ModeProduction
)

func NewMode(mode string) (Mode, error) {
	switch mode {
	case "production", "prod":
		return ModeProduction, nil
	case "development", "dev":
		return ModeDevelopment, nil
	default:
		return 0, errors.New("invalid mode provided")
	}
}
