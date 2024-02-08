package asynq

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type customAsynqLogger struct {
	service.ILogger
}

func (l customAsynqLogger) Debug(args ...interface{}) {
	l.ILogger.Debug("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Info(args ...interface{}) {
	l.ILogger.Info("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Warn(args ...interface{}) {
	l.ILogger.Warn("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Error(args ...interface{}) {
	l.ILogger.Error("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Fatal(args ...interface{}) {
	l.ILogger.Error("ASYNQ", "FATAL", args)
}
