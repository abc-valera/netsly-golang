package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type IJoke interface {
	GetByID(ctx context.Context, id string) (model.Joke, error)
	GetAllByUserID(ctx context.Context, userID string, s selector.Selector) (model.Jokes, error)
	SearchAllByTitle(ctx context.Context, keyword string, s selector.Selector) (model.Jokes, error)
}
