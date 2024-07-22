package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type like struct {
	executor boil.ContextExecutor
}

func NewLike(executor boil.ContextExecutor) command.ILike {
	return &like{
		executor: executor,
	}
}

func (l like) Create(ctx context.Context, req command.LikeCreateRequest) (model.Like, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	like := sqlboiler.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: req.Like.CreatedAt,
		DeletedAt: req.Like.DeletedAt,
	}
	err := like.Insert(ctx, l.executor, boil.Infer())
	return boilerSqliteDto.NewDomainLike(&like), boilerSqliteErrutil.HandleErr(err)
}

func (l like) Delete(ctx context.Context, userID string, jokeID string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	like, err := sqlboiler.FindLike(ctx, l.executor, userID, jokeID)
	if err != nil {
		return err
	}
	_, err = like.Delete(ctx, l.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
