package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestJokeCreate(t *testing.T) {
	ctx, r, entities := NewTest(t)

	user, err := entities.User.Create(ctx, entity.UserCreateRequest{
		Username: "testUsername",
		Email:    "testEmail@gmail.com",
		Password: "testPassword",
		Fullname: "testFullname",
		Status:   "testStatus",
	})
	r.NoError(err)
	r.NotEmpty(user)

	expected, err := entities.Joke.Create(ctx, entity.JokeCreateRequest{
		Title:       "testTitle",
		Text:        "testText",
		Explanation: "testExplanation",
		UserID:      user.ID,
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.Joke.GetByID(ctx, expected.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(expected, actual)
}

func TestJokeUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestJokeDelete(t *testing.T) {
	t.Skip("unimplemented")
}
