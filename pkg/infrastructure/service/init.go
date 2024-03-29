package service

import (
	"os"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/passwordMaker"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/taskQueuer/dummyTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/tokenMaker"
)

func Init() domain.Services {
	// Get all the environmental variables
	emailSenderServiceEnv := os.Getenv("EMAIL_SENDER_SERVICE")
	if emailSenderServiceEnv == "" {
		coderr.Fatal("EMAIL_SENDER_SERVICE environmental variable is not set")
	}
	taskQueuerServiceEnv := os.Getenv("TASK_QUEUER_SERVICE")
	if taskQueuerServiceEnv == "" {
		coderr.Fatal("TASK_QUEUER_SERVICE environmental variable is not set")
	}

	accessTokenDurationEnv := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessTokenDurationEnv == "" {
		coderr.Fatal("ACCESS_TOKEN_DURATION environmental variable is not set")
	}
	refreshTokenDurationEnv := os.Getenv("REFRESH_TOKEN_DURATION")
	if refreshTokenDurationEnv == "" {
		coderr.Fatal("REFRESH_TOKEN_DURATION environmental variable is not set")
	}
	signKeyEnv := os.Getenv("JWT_SIGN_KEY")
	if signKeyEnv == "" {
		coderr.Fatal("JWT_SIGN_KEY environmental variable is not set")
	}

	passwordMaker := passwordMaker.New()

	tokenMaker := tokenMaker.NewJWT(
		coderr.Must(time.ParseDuration(accessTokenDurationEnv)),
		coderr.Must(time.ParseDuration(refreshTokenDurationEnv)),
		signKeyEnv,
	)

	var emailSender service.IEmailSender
	switch emailSenderServiceEnv {
	case "dummy":
		emailSender = dummyEmailSender.New()
	default:
		coderr.Fatal("EMAIL_SENDER_SERVICE environmental variable is invalid")
	}

	var taskQueuer service.ITaskQueuer
	switch taskQueuerServiceEnv {
	case "dummy":
		taskQueuer = dummyTaskQueuer.New(emailSender)
	default:
		coderr.Fatal("TASK_QUEUER_SERVICE environmental variable is invalid")
	}

	return domain.NewServices(
		emailSender,
		passwordMaker,
		tokenMaker,
		taskQueuer,
	)
}
