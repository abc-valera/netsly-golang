package messaging

import (
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
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
	log.Fatalln("ASYNQ", "FATAL", args)
}
