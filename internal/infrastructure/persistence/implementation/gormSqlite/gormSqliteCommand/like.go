package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type like struct {
	db *gorm.DB
}

func NewLike(db *gorm.DB) command.ILike {
	return &like{
		db: db,
	}
}

func (c like) Create(ctx context.Context, req command.LikeCreateRequest) (model.Like, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	like := gormSqliteDto.Like{
		CreatedAt: req.Like.CreatedAt,
		DeletedAt: req.Like.DeletedAt,
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	}
	res := c.db.Create(&like)
	return gormSqliteDto.NewDomainLike(like), gormSqliteErrutil.HandleCommandResult(res)
}

func (c like) Delete(ctx context.Context, userID, jokeID string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	like := gormSqliteDto.Like{
		UserID: userID,
		JokeID: jokeID,
	}
	res := c.db.Delete(&like)
	return gormSqliteErrutil.HandleCommandResult(res)
}
