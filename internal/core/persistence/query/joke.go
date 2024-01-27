package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query/spec"
)

type IJoke interface {
	GetByID(ctx context.Context, id string) (model.Joke, error)
	GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Jokes, error)
	SearchByTitle(ctx context.Context, keyword string, params spec.SelectParams) (model.Jokes, error)
}
