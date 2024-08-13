package taskQueuerDummy

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type dummyTaskQueuer struct {
	emailSender service.IEmailSender
}

func New(emailSender service.IEmailSender) service.ITaskQueuer {
	return &dummyTaskQueuer{
		emailSender: emailSender,
	}
}

func (s dummyTaskQueuer) SendEmailTask(ctx context.Context, priority service.TaskPriority, email service.Email) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	// Some logic to enqueue the send email task
	time.Sleep(10 * time.Millisecond)
	span.AddEvent("Finished enqueuing the task")

	go func() {
		s.emailSender.SendEmail(email)
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
