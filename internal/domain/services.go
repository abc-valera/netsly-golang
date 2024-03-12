package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/logger"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger        logger.ILogger
	UUUIDMaker    service.IUuidMaker
	Time          service.ITimeMaker
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	TaskQueuer    service.ITaskQueuer
}

func NewServices(
	logger logger.ILogger,
	uuidMaker service.IUuidMaker,
	timer service.ITimeMaker,
	emailSender service.IEmailSender,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	taskQueue service.ITaskQueuer,
) Services {
	return Services{
		Logger:        logger,
		UUUIDMaker:    uuidMaker,
		Time:          timer,
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		TaskQueuer:    taskQueue,
	}
}
