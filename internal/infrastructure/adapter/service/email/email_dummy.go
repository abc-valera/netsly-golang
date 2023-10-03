package email

import "github.com/abc-valera/flugo-api-golang/internal/domain/service"

type dummyEmailSender struct{}

func NewDummyEmailSender() service.EmailSender {
	return &dummyEmailSender{}
}

func (d *dummyEmailSender) SendEmail(e service.Email) error {
	return nil
}
