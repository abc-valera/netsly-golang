package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	TaskQueuer    service.ITaskQueuer
}

func NewServices(
	emailSender service.IEmailSender,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	taskQueue service.ITaskQueuer,
) Services {
	return Services{
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		TaskQueuer:    taskQueue,
	}
}
