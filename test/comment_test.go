package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestCommentCreate(t *testing.T) {
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

	expected, err := entities.Comment.Create(ctx, entity.CommentCreateRequest{
		Text: "testText",

		UserID: user.ID,
		JokeID: joke.ID,
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.Comment.GetByID(ctx, expected.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(expected, actual)
}

func TestCommentUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestCommentDelete(t *testing.T) {
	t.Skip("unimplemented")
}
