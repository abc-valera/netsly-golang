package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
)

type IJoke interface {
	GetByID(ctx context.Context, id string) (model.Joke, error)
	GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Jokes, error)
	SearchByTitle(ctx context.Context, keyword string, params spec.SelectParams) (model.Jokes, error)
}
