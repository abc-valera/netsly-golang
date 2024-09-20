package test

import "testing"

func TestPassworderHashPassword(t *testing.T) {
	_, r, entities := NewTest(t)

	testPassword := "testPassword"

	hash, err := entities.Passworder.HashPassword(testPassword)
	r.NoError(err)
	r.NotEmpty(hash)
	r.NotEqual(testPassword, hash)
}

func TestPassworderCheckPassword(t *testing.T) {
	_, r, entities := NewTest(t)

	t.Run("Success", func(t *testing.T) {
		testPassword := "testPassword"

		hash, err := entities.Passworder.HashPassword(testPassword)
		r.NoError(err)

		err = entities.Passworder.CheckPassword(testPassword, hash)
		r.NoError(err)
	})
	t.Run("Failure", func(t *testing.T) {
		hash, err := entities.Passworder.HashPassword("testPassword")
		r.NoError(err)

		err = entities.Passworder.CheckPassword("anotherTestPassword", hash)
		r.Error(err)
	})
}
