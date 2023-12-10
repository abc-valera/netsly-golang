package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IJokeCommand interface {
	Create(ctx context.Context, req model.Joke) error
	Update(ctx context.Context, id string, req model.JokeUpdate) error
	Delete(ctx context.Context, id string) error
}
