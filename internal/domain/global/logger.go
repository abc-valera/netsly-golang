package global

type ILogger interface {
	Trace(message string, vals ...any)
	Debug(message string, vals ...any)
	Info(message string, vals ...any)
	Warn(message string, vals ...any)
	Error(message string, vals ...any)
}
