package service_test

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/emailSender/dummyEmailSender"
)

func TestSendEmail(t *testing.T) {
	r := initTest(t)
	emailSender := dummyEmailSender.New()

	email := service.Email{
		Subject: "Test",
		Content: "Test",
		To:      []string{"test1", "test2"},
	}
	r.NoError(emailSender.SendEmail(email))
}
