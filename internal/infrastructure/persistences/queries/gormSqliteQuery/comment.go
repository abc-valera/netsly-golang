package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type comment struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) query.IComment {
	return &comment{
		db: db,
	}
}

func (q comment) GetByID(ctx context.Context, id string) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var comment gormSqliteDto.Comment
	res := q.db.Where("id = ?", id).First(&comment)
	return comment.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}

func (q comment) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Comments, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var comments gormSqliteDto.Comments
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("joke_id = ?", jokeID).
		Find(&comments)
	return comments.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}
