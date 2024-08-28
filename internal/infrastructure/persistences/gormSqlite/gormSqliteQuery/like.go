package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type like struct {
	db *gorm.DB
}

func NewLike(db *gorm.DB) query.ILike {
	return &like{
		db: db,
	}
}

func (q like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var count int64
	res := q.db.Model(&model.Like{}).Where("joke_id = ?", jokeID).Count(&count)
	return int(count), res.Error
}

func (q like) GetAllByJokeID(ctx context.Context, jokeID string, selector selector.Selector) (model.Likes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var likes gormSqliteDto.Likes
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("joke_id = ?", jokeID).
		Find(&likes)
	return gormSqliteDto.NewDomainLikes(likes), res.Error
}
