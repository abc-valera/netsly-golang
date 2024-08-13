package emailSenderDummy

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type dummyEmailSender struct{}

func New() service.IEmailSender {
	return &dummyEmailSender{}
}

func (d dummyEmailSender) SendEmail(e service.Email) error {
	// Sending email
	time.Sleep(500 * time.Millisecond)

	global.Log().Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
