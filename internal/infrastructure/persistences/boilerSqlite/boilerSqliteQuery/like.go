package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
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
	return int(count), boilerSqliteErrutil.HandleErr(err)
}

func (l like) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Likes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	likes, err := sqlboiler.Likes(sqlboiler.LikeWhere.JokeID.EQ(jokeID)).All(ctx, l.executor)
	return boilerSqliteDto.NewDomainLikes(likes), boilerSqliteErrutil.HandleErr(err)
}
