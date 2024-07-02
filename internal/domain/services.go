package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	Logger        service.ILogger
	FileManager   service.IFileManager
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TaskQueuer    service.ITaskQueuer
}
