package services

import (
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/env"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/emailer/emailerDummy"
)

func NewServices() service.Services {
	var services service.Services

	switch emailerService := env.Load("EMAIL_SENDER_SERVICE"); emailerService {
	case "dummy":
		services.Emailer = emailerDummy.New()
	default:
		coderr.Fatal("Invalid Email Sender implementation provided: " + emailerService)
	}

	// Check if the services are valid
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(services))

	return services
}
