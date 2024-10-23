package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
)

type ILike interface {
	queryGeneric.IGetOneGetMany[model.Like]
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
}
