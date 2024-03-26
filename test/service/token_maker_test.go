package service_test

import (
	"testing"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/service/tokenMaker"
	"github.com/stretchr/testify/require"
)

func TestCreateAccessToken(t *testing.T) {
	r := initTest(t)
	tokenMaker := tokenMaker.NewJWT(time.Hour, 24*time.Hour, "12345678901234567890123456789012")

	token, err := tokenMaker.CreateAccessToken("test_id")
	r.NoError(err)
	r.NotEmpty(token)
}

func TestCreateRefreshToken(t *testing.T) {
	r := initTest(t)
	tokenMaker := tokenMaker.NewJWT(time.Hour, 24*time.Hour, "12345678901234567890123456789012")

	token, err := tokenMaker.CreateRefreshToken("test_id")
	r.NoError(err)
	r.NotEmpty(token)
}

func TestVerifyToken(t *testing.T) {
	initTest(t)
	tokenMaker := tokenMaker.NewJWT(time.Hour, 24*time.Hour, "12345678901234567890123456789012")

	t.Run("Verify Access Token", func(t *testing.T) {
		r := require.New(t)

		token, err := tokenMaker.CreateAccessToken("test_id")
		r.NoError(err)
		r.NotEmpty(token)

		payload, err := tokenMaker.VerifyToken(token)
		r.NoError(err)
		r.Equal(false, payload.IsRefresh)
	})

	t.Run("Verify Refresh Token", func(t *testing.T) {
		r := require.New(t)

		token, err := tokenMaker.CreateRefreshToken("test_id")
		r.NoError(err)
		r.NotEmpty(token)

		payload, err := tokenMaker.VerifyToken(token)
		r.NoError(err)
		r.Equal(true, payload.IsRefresh)
	})
}
