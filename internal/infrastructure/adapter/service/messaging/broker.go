package messaging

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

type broker struct {
	client *asynq.Client
	server *asynq.Server
	log    service.Logger
}

func NewMessagingBroker(
	redisUrl, redisUser, redisPass string,
	emailSender service.EmailSender,
	log service.Logger,
) service.MessageBroker {
	// Redis connection options
	redisOpts := asynq.RedisClientOpt{
		Addr:     redisUrl,
		Username: redisUser,
		Password: redisPass,
	}

	return &broker{
		client: asynq.NewClient(redisOpts),
		server: newAsynqServer(redisOpts, emailSender, log),
		log:    log,
	}
}

func (b *broker) SendEmailTask(ctx context.Context, priority service.Priority, email service.Email) error {
	task, err := NewSendEmailTask("email")
	if err != nil {
		return err
	}

	info, err := b.client.Enqueue(
		task,
		asynq.MaxRetry(5),
		asynq.ProcessIn(2*time.Second),
		asynq.Queue(string(priority)),
	)
	if err != nil {
		return err
	}
	b.log.Info("ENQUEUED TASK",
		"type", task.Type(),
		"queue", info.Queue,
		"max_retry", info.MaxRetry,
	)
	return nil
}
