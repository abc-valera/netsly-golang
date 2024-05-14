package mode

// Mode is the application running mode
type Mode int

const (
	Development Mode = iota
	Production
)
