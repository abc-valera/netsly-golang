package dummy

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/core/service"
)

type broker struct {
	emailSender service.IEmailSender
}

func NewMessagingBroker(emailSender service.IEmailSender) service.IMessageBroker {
	return &broker{
		emailSender: emailSender,
	}
}

func (b broker) SendEmailTask(ctx context.Context, priority service.Priority, email service.Email) error {
	go func() {
		b.emailSender.SendEmail(email)
		global.Log.Info("SENT EMAIL",
			"to", email.To,
			"subject", email.Subject,
		)
	}()

	global.Log.Info("ENQUEUED TASK",
		"type", "email",
		"queue", string(priority),
		"max_retry", 5,
	)
	return nil
}
