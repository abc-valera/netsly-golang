package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
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

func (l like) Create(ctx context.Context, req model.Like) (model.Like, error) {
	like := sqlboiler.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: req.CreatedAt,
	}
	err := like.Insert(ctx, l.executor, boil.Infer())
	return boilerDto.NewDomainLikeWithErrHandle(&like, err)
}

func (l like) Delete(ctx context.Context, userID string, jokeID string) error {
	like, err := sqlboiler.FindLike(ctx, l.executor, userID, jokeID)
	if err != nil {
		return err
	}
	_, err = like.Delete(ctx, l.executor)
	return errutil.HandleErr(err)
}
