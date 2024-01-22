package service

import (
	"errors"
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
)

type Services struct {
	EmailSender   IEmailSender
	PasswordMaker IPasswordMaker
	TokenMaker    ITokenMaker
	MessageBroker IMessageBroker
}

func NewServices(
	logger ILogger,
	emailSender IEmailSender,
	passwordMaker IPasswordMaker,
	tokenMaker ITokenMaker,
	messageBroker IMessageBroker,
) Services {
	if logger == nil {
		log.Fatal(coderr.NewInternal(errors.New("logger is nil")))
	}
	if emailSender == nil {
		log.Fatal(coderr.NewInternal(errors.New("email sender is nil")))
	}
	if passwordMaker == nil {
		log.Fatal(coderr.NewInternal(errors.New("password maker is nil")))
	}
	if tokenMaker == nil {
		log.Fatal(coderr.NewInternal(errors.New("token maker is nil")))
	}
	if messageBroker == nil {
		log.Fatal(coderr.NewInternal(errors.New("message broker is nil")))
	}

	Log = logger

	return Services{
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		MessageBroker: messageBroker,
	}
}
