package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger        service.ILogger
	Validator     service.IValidator
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	MessageBroker service.IMessageBroker
}

func NewServices(
	logger service.ILogger,
	validator service.IValidator,
	emailSender service.IEmailSender,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	messageBroker service.IMessageBroker,
) Services {
	return Services{
		Logger:        logger,
		Validator:     validator,
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		MessageBroker: messageBroker,
	}
}
