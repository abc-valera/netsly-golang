package service

type ILogger interface {
	Debug(message string, vals ...interface{})
	Info(message string, vals ...interface{})
	Warn(message string, vals ...interface{})
	Error(message string, vals ...interface{})
}
