package entity

import (
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type IEmailer interface {
	SendEmail(e EmailSendRequest) error
}

func newEmailer(dep IDependency) IEmailer {
	return emailer{
		IDependency: dep,
	}
}

type emailer struct {
	IDependency
}

type EmailSendRequest struct {
	To      []string `validate:"required,dive,email"`
	Subject string   `validate:"required"`
	Content string   `validate:"required"`
}

func (e emailer) SendEmail(req EmailSendRequest) error {
	if err := global.Validate().Struct(req); err != nil {
		return err
	}

	return e.IDependency.S().Emailer.SendEmail(service.EmailSendRequest{
		To:      req.To,
		Subject: req.Subject,
		Content: req.Content,
	})
}
