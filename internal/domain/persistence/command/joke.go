package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

var ErrJokeAlreadyExists = coderr.NewCodeMessage(
	coderr.CodeAlreadyExists,
	"Joke with such title already exists by such user",
)

type IJoke interface {
	Create(ctx context.Context, userID string, req model.Joke) (model.Joke, error)
	Update(ctx context.Context, id string, req JokeUpdate) (model.Joke, error)
	Delete(ctx context.Context, id string) error
}

type JokeUpdate struct {
	Title       *string
	Text        *string
	Explanation *string
}
