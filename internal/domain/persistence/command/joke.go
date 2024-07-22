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
	Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error)
	Update(ctx context.Context, id string, req JokeUpdateRequest) (model.Joke, error)
	Delete(ctx context.Context, id string) error
}

type JokeCreateRequest struct {
	Joke   model.Joke
	UserID string
}

type JokeUpdateRequest struct {
	Title       *string
	Text        *string
	Explanation *string
}
