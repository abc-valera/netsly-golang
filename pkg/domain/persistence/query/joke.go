package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
)

type IJoke interface {
	GetByID(ctx context.Context, id string) (model.Joke, error)
	GetAllByUserID(ctx context.Context, userID string, params selector.Selector) (model.Jokes, error)
	SearchByTitle(ctx context.Context, keyword string, params selector.Selector) (model.Jokes, error)
}
