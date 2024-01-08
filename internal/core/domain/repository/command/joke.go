package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
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
