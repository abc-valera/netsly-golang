package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type like struct {
	executor boil.ContextExecutor
}

func NewLike(executor boil.ContextExecutor) query.ILike {
	return &like{
		executor: executor,
	}
}

func (l like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	count, err := sqlboiler.Likes(sqlboiler.LikeWhere.JokeID.EQ(jokeID)).Count(ctx, l.executor)
	return int(count), errutil.HandleErr(err)
}

func (l like) GatAllByJokeID(ctx context.Context, jokeID string) (model.Likes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	likes, err := sqlboiler.Likes(sqlboiler.LikeWhere.JokeID.EQ(jokeID)).All(ctx, l.executor)
	return boilerDto.NewDomainLikesWithErrHandle(likes, err)
}