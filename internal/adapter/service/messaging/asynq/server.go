package asynq

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/hibiken/asynq"
)

func newAsynqServer(
	redisOpts asynq.RedisClientOpt,
	emailSender service.IEmailSender,
	log service.ILogger,
) *asynq.Server {
	// Custom error handler
	errHandler := asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
		code := codeerr.ErrorCode(err)
		msg := codeerr.ErrorMessage(err)
		if code == "" {
			code = codeerr.CodeInternal
		}
		log.Error("PROCESS TASK",
			"code", code,
			"msg", msg,
			"error", err,
			"task", task.Type(),
			"payload:", string(task.Payload()))
	})

	server := asynq.NewServer(
		redisOpts,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				string(service.Critical): 6,
				string(service.Default):  3,
				string(service.Low):      1,
			},
			ErrorHandler:    errHandler,
			Logger:          &customAsynqLogger{log},
			ShutdownTimeout: time.Millisecond * 100,
		},
	)

	// Allocating tasks
	mux := asynq.NewServeMux()
	mux.Handle(typeSendEmail, newSendEmailProcessor(emailSender, log))

	// Running server
	server.Start(mux)

	return server
}
