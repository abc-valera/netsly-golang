package global

// Mode is the application running mode
type Mode int

const (
	ModeDevelopment Mode = iota + 1
	ModeProduction
)
