package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger      service.ILogger
	EmailSender service.IEmailSender
	Passworder  service.IPassworder
	TaskQueuer  service.ITaskQueuer
}
