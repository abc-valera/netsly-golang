package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteQuery/gormSelector"
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
	return gormSqliteDto.NewDomainComment(comment), gormSqliteErrutil.HandleQueryResult(res)
}

func (q comment) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Comments, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var comments gormSqliteDto.Comments
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("joke_id = ?", jokeID).
		Find(&comments)
	return gormSqliteDto.NewDomainComments(comments), gormSqliteErrutil.HandleQueryResult(res)
}
