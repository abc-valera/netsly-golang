package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel/trace"
)

type like struct {
	db bun.IDB
	queryGeneric.IGetOneGetMany[model.Like]
}

func NewLike(db bun.IDB) query.ILike {
	return &like{
		db: db,
		IGetOneGetMany: bunSqliteQueryGeneric.New(
			db,
			bunSqliteDto.NewLike,
		),
	}
}

func (q like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	count, err := q.db.NewSelect().
		Model((*bunSqliteDto.Like)(nil)).
		Where("joke_id = ?", jokeID).
		Count(ctx)
	return int(count), bunSqliteErrors.HandleQueryResult(err)
}
