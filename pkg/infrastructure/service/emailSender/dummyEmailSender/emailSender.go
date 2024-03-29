package dummyEmailSender

import (
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
)

type dummyEmailSender struct {
}

func New() service.IEmailSender {
	return &dummyEmailSender{}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	global.Log().Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
