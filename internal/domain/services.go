package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/logger"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger        logger.ILogger
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	TaskQueuer    service.ITaskQueuer
}

func NewServices(
	logger logger.ILogger,
	emailSender service.IEmailSender,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	messageBroker service.ITaskQueuer,
) Services {
	return Services{
		Logger:        logger,
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		TaskQueuer:    messageBroker,
	}
}
