package messaging

import (
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type customAsynqLogger struct {
	service.Logger
}

func (l customAsynqLogger) Debug(args ...interface{}) {
	l.Logger.Debug("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Info(args ...interface{}) {
	l.Logger.Info("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Warn(args ...interface{}) {
	l.Logger.Warn("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Error(args ...interface{}) {
	l.Logger.Error("ASYNQ", "msg", args)
}

func (l customAsynqLogger) Fatal(args ...interface{}) {
	log.Fatalln("ASYNQ", "FATAL", args)
}
