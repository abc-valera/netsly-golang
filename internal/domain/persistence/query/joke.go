package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

type IJoke interface {
	GetByID(ctx context.Context, id string) (model.Joke, error)
	GetAllByUserID(ctx context.Context, userID string, selector selector.Selector) (model.Jokes, error)
	SearchAllByTitle(ctx context.Context, keyword string, selector selector.Selector) (model.Jokes, error)
}
