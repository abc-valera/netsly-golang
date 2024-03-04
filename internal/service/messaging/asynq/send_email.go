package asynq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	typeSendEmail = "email:send"
)

type sendEmailPayload struct {
	Email string
}

func NewSendEmailTask(email string) (*asynq.Task, error) {
	payload, err := json.Marshal(sendEmailPayload{Email: email})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(typeSendEmail, payload), nil
}

type sendEmailProcessor struct {
	emailSender service.IEmailSender
}

func newSendEmailProcessor(
	emailSender service.IEmailSender,
) *sendEmailProcessor {
	return &sendEmailProcessor{
		emailSender: emailSender,
	}
}

func (p sendEmailProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload sendEmailPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	global.Log().Info("TASK",
		"type", task.Type())
	return p.emailSender.SendEmail(service.Email{})
}