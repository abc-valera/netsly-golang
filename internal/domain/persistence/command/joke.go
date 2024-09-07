package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrJokeAlreadyExists = coderr.NewCodeMessage(
	coderr.CodeAlreadyExists,
	"Joke with such title already exists by such user",
)

type IJoke interface {
	Create(ctx context.Context, req model.Joke) (model.Joke, error)
	Update(ctx context.Context, ids model.Joke, req JokeUpdateRequest) (model.Joke, error)
	Delete(ctx context.Context, req model.Joke) error
}

type JokeUpdateRequest struct {
	UpdatedAt time.Time

	Title       *string
	Text        *string
	Explanation *string
}
