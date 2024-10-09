package emailerDummy

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type dummyEmailer struct{}

func New() service.IEmailer {
	return &dummyEmailer{}
}

func (dummyEmailer) SendEmail(e service.EmailSendRequest) error {
	// Sending email
	time.Sleep(500 * time.Millisecond)

	global.Log().Info("EMAIL_SENT", "to", e.To, "subject", e.Subject, "body", e.Content)
	return nil
}
