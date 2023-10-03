package service

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/service/email"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/service/logger"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/service/messaging"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/service/password"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/service/token"
)

func NewServices(
	accessTokenDuration, refreshTokenDuration time.Duration,
	redisUrl, redisUser, redisPass string,
) (service.Services, error) {
	emailSender := email.NewDummyEmailSender()
	passwordMaker := password.NewPasswordMaker()
	tokenMaker := token.NewTokenMaker(accessTokenDuration, refreshTokenDuration)
	logger := logger.NewSlogLogger()
	messageBroker := messaging.NewMessagingBroker(redisUrl, redisUser, redisPass, emailSender, logger)

	return service.Services{
		EmailSender:    emailSender,
		PassswordMaker: passwordMaker,
		TokenMaker:     tokenMaker,
		Logger:         logger,
		MessageBroker:  messageBroker,
	}, nil
}
