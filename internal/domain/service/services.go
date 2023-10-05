package service

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
)

type Services struct {
	EmailSender   EmailSender
	PasswordMaker PasswordMaker
	TokenMaker    TokenMaker
	Logger        Logger
	MessageBroker MessageBroker
}

func NewServices(
	emailSender EmailSender,
	passwordMaker PasswordMaker,
	tokenMaker TokenMaker,
	logger Logger,
	messageBroker MessageBroker,
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
	if logger == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("logger is nil"))
	}
	if messageBroker == nil {
		return Services{}, codeerr.NewInternal("NewServices", errors.New("messageBroker is nil"))
	}
	return Services{
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		Logger:        logger,
		MessageBroker: messageBroker,
	}, nil
}
