package logger

type ILogger interface {
	Trace(message string, vals ...interface{})
	Debug(message string, vals ...interface{})
	Info(message string, vals ...interface{})
	Warn(message string, vals ...interface{})
	Error(message string, vals ...interface{})
	Fatal(message string, vals ...interface{})
}
