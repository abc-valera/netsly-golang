package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger        service.ILogger
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	TaskQueuer    service.ITaskQueuer
}
