package asynq

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

func newAsynqServer(
	redisOpts asynq.RedisClientOpt,
	emailSender service.IEmailSender,
) *asynq.Server {
	// Custom error handler
	errHandler := asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
		code := coderr.ErrorCode(err)
		msg := coderr.ErrorMessage(err)
		if code == "" {
			code = coderr.CodeInternal
		}
		global.Log().Error("PROCESS TASK",
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
			Logger:          &customAsynqLogger{global.Log()},
			ShutdownTimeout: time.Millisecond * 100,
		},
	)

	// Allocating tasks
	mux := asynq.NewServeMux()
	mux.Handle(typeSendEmail, newSendEmailProcessor(emailSender))

	// Running server
	server.Start(mux)

	return server
}
