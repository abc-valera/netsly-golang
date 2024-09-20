package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestEmailerSendEmail(t *testing.T) {
	_, r, entities := NewTest(t)

	err := entities.Emailer.SendEmail(entity.EmailSendRequest{
		To:      []string{"test@netsly-test.com"},
		Subject: "testSubject",
		Content: "testContent",
	})
	r.NoError(err)
}
