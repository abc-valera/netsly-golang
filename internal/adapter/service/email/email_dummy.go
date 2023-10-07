package email

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type dummyEmailSender struct {
	log service.Logger
}

func NewDummyEmailSender(log service.Logger) service.EmailSender {
	return &dummyEmailSender{
		log: log,
	}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	d.log.Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
