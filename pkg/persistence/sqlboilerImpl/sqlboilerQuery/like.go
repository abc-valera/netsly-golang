package sqlboilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type like struct {
	executor boil.ContextExecutor
}

func newLike(executor boil.ContextExecutor) query.ILike {
	return &like{
		executor: executor,
	}
}

func (l like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	count, err := sqlboiler.Likes(sqlboiler.LikeWhere.JokeID.EQ(jokeID)).Count(ctx, l.executor)
	return int(count), errors.HandleErr(err)
}
