package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type ILike interface {
	GatAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Likes, error)
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
}
