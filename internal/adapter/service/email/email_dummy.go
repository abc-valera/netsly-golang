package email

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type dummyEmailSender struct {
}

func NewDummyEmailSender() service.IEmailSender {
	return &dummyEmailSender{}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	service.Log.Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
