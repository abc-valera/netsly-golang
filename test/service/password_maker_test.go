package service_test

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/passwordMaker"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	r := initTest(t)
	passwordMaker := passwordMaker.New()

	hashedPassword, err := passwordMaker.HashPassword("test_password")
	r.NoError(err)
	r.NotEmpty(hashedPassword)
}

func TestCheckPassword(t *testing.T) {
	initTest(t)
	passwordMaker := passwordMaker.New()

	t.Run("Success", func(t *testing.T) {
		r := require.New(t)

		password := "test_password"

		hashedPassword, err := passwordMaker.HashPassword(password)
		r.NoError(err)
		r.NotEmpty(hashedPassword)

		r.NoError(passwordMaker.CheckPassword(password, hashedPassword))
	})

	t.Run("Provided password doesn't match", func(t *testing.T) {
		r := require.New(t)

		hashedPassword, err := passwordMaker.HashPassword("test_password")
		r.NoError(err)
		r.NotEmpty(hashedPassword)

		r.Equal(passwordMaker.CheckPassword("incorrect_password", hashedPassword), service.ErrPasswordDontMatch)
	})
}
