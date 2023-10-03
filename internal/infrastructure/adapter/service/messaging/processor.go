package messaging

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

const (
	queueCritical = "critical"
	queueDefault  = "default"
)

type redisTaskProcessor struct {
	server      *asynq.Server
	emailSender service.EmailSender
	log         service.Logger
}

func newRedisTaskProcessor(
	redisOpt *asynq.RedisClientOpt,
	emailSender service.EmailSender,
	log service.Logger,
) *redisTaskProcessor {
	return &redisTaskProcessor{
		server: asynq.NewServer(
			redisOpt,
			asynq.Config{
				Queues: map[string]int{
					queueCritical: 10,
					queueDefault:  5,
				},
				ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
					log.Error("PROCESS TASK",
						"code", codeerr.ErrorCode(err),
						"msg", codeerr.ErrorMessage(err),
						"error", err,
						"task", task.Type(),
						"payload:", string(task.Payload()))
				}),
			},
		),
		emailSender: emailSender,
		log:         log,
	}
}

func (p *redisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(taskSendEmail, p.ProccessTaskSendEmail)

	return p.server.Start(mux)
}
