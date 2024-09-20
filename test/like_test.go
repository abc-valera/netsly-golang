package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestLikeCreate(t *testing.T) {
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

	joke, err := entities.Joke.Create(ctx, entity.JokeCreateRequest{
		Title:       "testTitle",
		Text:        "testText",
		Explanation: "testExplanation",
		UserID:      user.ID,
	})
	r.NoError(err)
	r.NotEmpty(joke)

	expected, err := entities.Like.Create(ctx, entity.LikeCreateRequest{
		UserID: user.ID,
		JokeID: joke.ID,
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.Like.CountByJokeID(ctx, joke.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(1, actual)
}

func TestLikeDelete(t *testing.T) {
	t.Skip("unimplemented")
}
