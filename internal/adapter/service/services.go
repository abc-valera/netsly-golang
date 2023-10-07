package service

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/messaging"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

func NewServices(
	accessTokenDuration, refreshTokenDuration time.Duration,
	redisUrl, redisUser, redisPass string,
) (service.Services, error) {
	logger := logger.NewSlogLogger()
	emailSender := email.NewDummyEmailSender(logger)
	passwordMaker := password.NewPasswordMaker()
	tokenMaker := token.NewTokenMaker(accessTokenDuration, refreshTokenDuration)
	messageBroker := messaging.NewMessagingBroker(redisUrl, redisUser, redisPass, emailSender, logger)

	return service.NewServices(
		emailSender,
		passwordMaker,
		tokenMaker,
		logger,
		messageBroker,
	)
}
