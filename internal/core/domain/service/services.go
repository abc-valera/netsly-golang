package service

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

type Services struct {
	EmailSender   IEmailSender
	PasswordMaker IPasswordMaker
	TokenMaker    ITokenMaker
	MessageBroker IMessageBroker
}

func NewServices(
	emailSender IEmailSender,
	passwordMaker IPasswordMaker,
	tokenMaker ITokenMaker,
	messageBroker IMessageBroker,
) (Services, error) {
	if emailSender == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("emailSender is nil"))
	}
	if passwordMaker == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("passwordMaker is nil"))
	}
	if tokenMaker == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("tokenMaker is nil"))
	}
	if messageBroker == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("messageBroker is nil"))
	}
	return Services{
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		MessageBroker: messageBroker,
	}, nil
}
