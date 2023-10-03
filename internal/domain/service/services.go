package service

type Services struct {
	EmailSender    EmailSender
	PassswordMaker PasswordMaker
	TokenMaker     TokenMaker
	Logger         Logger
	MessageBroker  MessageBroker
}
