package email

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type dummyEmailSender struct {
	log service.ILogger
}

func NewDummyEmailSender(log service.ILogger) service.IEmailSender {
	return &dummyEmailSender{
		log: log,
	}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	d.log.Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
