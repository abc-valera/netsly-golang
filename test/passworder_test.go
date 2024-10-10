package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassworderHashPassword(t *testing.T) {
	_, r, entities := NewTest(t)

	testPassword := "testPassword"

	hash, err := entities.Passworder.HashPassword(testPassword)
	r.NoError(err)
	r.NotEmpty(hash)
	r.NotEqual(testPassword, hash)
}

func TestPassworderCheckPassword(t *testing.T) {
	_, _, entities := NewTest(t)

	t.Run("Success", func(t *testing.T) {
		testPassword := "testPassword"

		hash, err := entities.Passworder.HashPassword(testPassword)
		require.NoError(t, err)

		err = entities.Passworder.CheckPassword(testPassword, hash)
		require.NoError(t, err)
	})
	t.Run("Failure", func(t *testing.T) {
		hash, err := entities.Passworder.HashPassword("testPassword")
		require.NoError(t, err)

		err = entities.Passworder.CheckPassword("anotherTestPassword", hash)
		require.Error(t, err)
	})
}
