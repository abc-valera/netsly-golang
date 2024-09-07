package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type comment struct {
	db bun.IDB
}

func NewComment(db bun.IDB) query.IComment {
	return &comment{
		db: db,
	}
}

func (q comment) GetByID(ctx context.Context, id string) (model.Comment, error) {
	comment := bunSqliteDto.Comment{}
	err := q.db.NewSelect().Model(&comment).Where("id = ?", id).Scan(ctx)
	return comment.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q comment) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Comments, error) {
	comments := bunSqliteDto.Comments{}
	err := q.db.NewSelect().Model(&comments).Where("joke_id = ?", jokeID).Scan(ctx)
	return comments.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}
