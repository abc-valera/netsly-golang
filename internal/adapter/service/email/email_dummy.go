package email

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type dummyEmailSender struct {
}

func NewDummyEmailSender() service.IEmailSender {
	return &dummyEmailSender{}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	global.Log.Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
