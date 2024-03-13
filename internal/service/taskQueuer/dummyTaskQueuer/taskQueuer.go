package dummyTaskQueuer

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type dummyTaskQueuer struct {
	emailSender service.IEmailSender
}

func New(emailSender service.IEmailSender) service.ITaskQueuer {
	return &dummyTaskQueuer{
		emailSender: emailSender,
	}
}

func (b dummyTaskQueuer) SendEmailTask(ctx context.Context, priority service.TaskPriority, email service.Email) error {
	go func() {
		b.emailSender.SendEmail(email)
		global.Log().Info("SENT EMAIL",
			"to", email.To,
			"subject", email.Subject,
		)
	}()

	global.Log().Info("ENQUEUED TASK",
		"type", "email",
		"queue", string(priority),
		"max_retry", 5,
	)
	return nil
}
