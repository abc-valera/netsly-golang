package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

var (
	ErrJokeOwnerTitleAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Joke with such title already exists by such user")
)

type IJoke interface {
	Create(ctx context.Context, req model.Joke) error
	Update(ctx context.Context, id string, req JokeUpdate) error
	Delete(ctx context.Context, id string) error
}

type JokeUpdate struct {
	Title       *string
	Text        *string
	Explanation *string
}
