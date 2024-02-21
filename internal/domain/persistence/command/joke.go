package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

var (
	ErrJokeOwnerTitleAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Joke with such title already exists by such user")
)

type IJoke interface {
	Create(ctx context.Context, req model.Joke) (model.Joke, error)
	Update(ctx context.Context, id string, req JokeUpdate) (model.Joke, error)
	Delete(ctx context.Context, id string) error
}

type JokeUpdate struct {
	Title       *string
	Text        *string
	Explanation *string
}
