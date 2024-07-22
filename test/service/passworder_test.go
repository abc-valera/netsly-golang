package service_test

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/passworder"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	r := initTest(t)
	passworder := passworder.New()

	hashedPassword, err := passworder.HashPassword("test_password")
	r.NoError(err)
	r.NotEmpty(hashedPassword)
}

func TestCheckPassword(t *testing.T) {
	initTest(t)
	passworder := passworder.New()

	t.Run("Success", func(t *testing.T) {
		r := require.New(t)

		password := "test_password"

		hashedPassword, err := passworder.HashPassword(password)
		r.NoError(err)
		r.NotEmpty(hashedPassword)

		r.NoError(passworder.CheckPassword(password, hashedPassword))
	})

	t.Run("Provided password doesn't match", func(t *testing.T) {
		r := require.New(t)

		hashedPassword, err := passworder.HashPassword("test_password")
		r.NoError(err)
		r.NotEmpty(hashedPassword)

		r.Equal(passworder.CheckPassword("incorrect_password", hashedPassword), service.ErrPasswordDontMatch)
	})
}
