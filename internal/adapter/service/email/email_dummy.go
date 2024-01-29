package email

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/core/service"
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
