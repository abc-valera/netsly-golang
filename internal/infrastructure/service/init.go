package service

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/nopLogger"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/passwordMaker"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/taskQueuer/dummyTaskQueuer"
)

func New(
	loggerService string,

	emailSenderService string,

	taskQueuerService string,
) domain.Services {
	var services domain.Services

	services.PasswordMaker = passwordMaker.New()

	switch loggerService {
	case "nop":
		services.Logger = nopLogger.New()
	case "slog":
		services.Logger = slogLogger.New()
	default:
		coderr.Fatal("Invalid logger implementation provided. Should be 'nop' or 'slog'")
	}

	switch emailSenderService {
	case "dummy":
		services.EmailSender = dummyEmailSender.New()
	default:
		coderr.Fatal("Invalid email sender implementation provided. Should be 'dummy'")
	}

	switch taskQueuerService {
	case "dummy":
		services.TaskQueuer = dummyTaskQueuer.New(services.EmailSender)
	default:
		coderr.Fatal("Invalid task queuer implementation provided. Should be 'dummy'")
	}

	return services
}
