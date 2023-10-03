package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

// Name of the task
const taskSendEmail = "task:send_email"

type payloadSendEmail struct {
	Email service.Email `json:"email"`
}

func (d *redisTaskDistributor) DistributeTaskSendEmail(
	ctx context.Context,
	payload payloadSendEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(&payload)
	if err != nil {
		return codeerr.NewInternal("DistributeTaskSendEmail", err)
	}

	task := asynq.NewTask(taskSendEmail, jsonPayload, opts...)
	info, err := d.client.EnqueueContext(ctx, task)
	if err != nil {
		return codeerr.NewInternal("DistributeTaskSendEmail", err)
	}

	d.log.Info("ENQUEUED TASK",
		"type", task.Type(),
		"queue", info.Queue,
		"max_retry", info.MaxRetry,
	)

	return nil
}

// TODO: error handling with asynq.SkipRetry
func (p *redisTaskProcessor) ProccessTaskSendEmail(c context.Context, task *asynq.Task) error {
	var payload payloadSendEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("%w %w", codeerr.NewInternal("ProccessTaskSendVerifyEmail", err), asynq.SkipRetry)
	}

	p.emailSender.SendEmail(payload.Email)

	p.log.Info("PROCESSED TASK",
		"type", task.Type(),
		"email", payload.Email,
	)

	return nil
}
