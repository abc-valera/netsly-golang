package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

type messagingBroker struct {
	distributor *redisTaskDistributor
	processor   *redisTaskProcessor
}

func NewMessagingBroker(
	redisUrl, redisUser, redisPass string,
	emailSender service.EmailSender,
	log service.Logger,
) service.MessageBroker {
	redisOpt := &asynq.RedisClientOpt{
		Addr:     redisUrl,
		Username: redisUser,
		Password: redisPass,
	}

	proc := newRedisTaskProcessor(redisOpt, emailSender, log)
	proc.log.Info("Starting task processor")
	go proc.Start()

	return &messagingBroker{
		distributor: newRedisTaskDistributor(redisOpt, log),
		processor:   proc,
	}
}

func (m *messagingBroker) SendEmailTask(ctx context.Context, priority service.Priority, email service.Email) error {
	taskPayload := payloadSendEmail{
		Email: email,
	}

	opts := []asynq.Option{
		asynq.MaxRetry(5),
		asynq.ProcessIn(2 * time.Second),
		asynq.Queue(fmt.Sprint(priority)),
	}

	return m.distributor.DistributeTaskSendEmail(ctx, taskPayload, opts...)
}
