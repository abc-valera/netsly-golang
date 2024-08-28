package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type like struct {
	db bun.IDB
}

func NewLike(db bun.IDB) command.ILike {
	return &like{
		db: db,
	}
}

func (c like) Create(ctx context.Context, req command.LikeCreateRequest) (model.Like, error) {
	like := bunSqliteDto.Like{
		CreatedAt: req.Like.CreatedAt,
		DeletedAt: req.Like.DeletedAt,

		UserID: req.UserID,
		JokeID: req.JokeID,
	}

	res, err := c.db.NewInsert().Model(&like).Exec(ctx)
	return like.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c like) Delete(ctx context.Context, userID string, jokeID string) error {
	like := bunSqliteDto.Like{
		UserID: userID,
		JokeID: jokeID,
	}
	res, err := c.db.NewDelete().Model(&like).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
