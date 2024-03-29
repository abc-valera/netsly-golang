package service_test

import (
	"context"
	"testing"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/abc-valera/netsly-api-golang/pkg/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/pkg/service/taskQueuer/dummyTaskQueuer"
)

func TestSendEmailTask(t *testing.T) {
	r := initTest(t)
	ctx := context.Background()
	taskQueuer := dummyTaskQueuer.New(dummyEmailSender.New())

	email := service.Email{
		Subject: "Test",
		Content: "Test",
		To:      []string{"test", "test"},
	}
	r.NoError(taskQueuer.SendEmailTask(ctx, service.Default, email))
}
