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

type joke struct {
	db *gorm.DB
}

func NewJoke(db *gorm.DB) query.IJoke {
	return &joke{
		db: db,
	}
}

func (q joke) GetByID(ctx context.Context, id string) (model.Joke, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var joke gormSqliteDto.Joke
	res := q.db.Where("id = ?", id).First(&joke)
	return gormSqliteDto.NewDomainJoke(joke), gormSqliteErrutil.HandleQueryResult(res)
}

func (q joke) SearchAllByTitle(ctx context.Context, keyword string, selector selector.Selector) (model.Jokes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var jokes gormSqliteDto.Jokes
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("title LIKE ?", "%"+keyword+"%").
		Find(&jokes)
	return gormSqliteDto.NewDomainJokes(jokes), gormSqliteErrutil.HandleQueryResult(res)
}

func (q joke) GetAllByUserID(ctx context.Context, userID string, selector selector.Selector) (model.Jokes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var jokes gormSqliteDto.Jokes
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&jokes)
	return gormSqliteDto.NewDomainJokes(jokes), gormSqliteErrutil.HandleQueryResult(res)
}
