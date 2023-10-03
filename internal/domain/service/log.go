package service

type Logger interface {
	Debug(msg string, vals ...interface{})
	Info(msg string, vals ...interface{})
	Warn(msg string, vals ...interface{})
	Error(msg string, vals ...interface{})
}
