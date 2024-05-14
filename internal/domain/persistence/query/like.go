package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type ILike interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
	GetAllByJokeID(ctx context.Context, jokeID string) (model.Likes, error)
}
