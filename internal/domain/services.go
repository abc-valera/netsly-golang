package domain

import (
	"errors"
	"log"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Services struct {
	EmailSender   service.IEmailSender
	PasswordMaker service.IPasswordMaker
	TokenMaker    service.ITokenMaker
	MessageBroker service.IMessageBroker
}

func NewServices(
	logger service.ILogger,
	emailSender service.IEmailSender,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	messageBroker service.IMessageBroker,
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

	global.Log = logger

	return Services{
		EmailSender:   emailSender,
		PasswordMaker: passwordMaker,
		TokenMaker:    tokenMaker,
		MessageBroker: messageBroker,
	}
}
