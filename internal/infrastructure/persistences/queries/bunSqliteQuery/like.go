package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
)

type like struct {
	db bun.IDB
}

func NewLike(db bun.IDB) query.ILike {
	return &like{
		db: db,
	}
}

func (q like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	count, err := q.db.NewSelect().Model((*bunSqliteDto.Like)(nil)).Where("joke_id = ?", jokeID).Count(ctx)
	return int(count), bunSqliteErrors.HandleQueryResult(err)
}

func (q like) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Likes, error) {
	likes := bunSqliteDto.Likes{}
	err := q.db.NewSelect().Model(&likes).Where("joke_id = ?", jokeID).Scan(ctx)
	return likes.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}
