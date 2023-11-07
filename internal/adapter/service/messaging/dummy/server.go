package dummy

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type broker struct {
	emailSender service.IEmailSender
	log         service.ILogger
}

func NewMessagingBroker(emailSender service.IEmailSender, log service.ILogger) service.IMessageBroker {
	return &broker{
		emailSender: emailSender,
		log:         log,
	}
}

func (b broker) SendEmailTask(ctx context.Context, priority service.Priority, email service.Email) error {
	go func() {
		b.emailSender.SendEmail(email)
		b.log.Info("SENT EMAIL",
			"to", email.To,
			"subject", email.Subject,
		)
	}()

	b.log.Info("ENQUEUED TASK",
		"type", "email",
		"queue", string(priority),
		"max_retry", 5,
	)
	return nil
}
