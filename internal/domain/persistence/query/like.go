package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type ILike interface {
	GatAllByJokeID(ctx context.Context, jokeID string) (model.Likes, error)
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
}
