package service

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/messaging/dummy"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

func NewServices(
	accessTokenDuration, refreshTokenDuration time.Duration,
	redisUrl, redisUser, redisPass string,
) (service.Services, error) {
	service.Log = logger.NewSlogLogger()

	emailSender := email.NewDummyEmailSender()

	passwordMaker := password.NewPasswordMaker()

	tokenMaker := token.NewTokenMaker(accessTokenDuration, refreshTokenDuration)

	messageBroker := dummy.NewMessagingBroker(emailSender)

	return service.NewServices(
		emailSender,
		passwordMaker,
		tokenMaker,
		messageBroker,
	)
}
