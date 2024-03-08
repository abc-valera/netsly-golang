package sqlboilercommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type like struct {
	executor boil.ContextExecutor
}

func newLike(executor boil.ContextExecutor) command.ILike {
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
	return dto.ToDomainLikeWithErrHandle(&like, err)
}

func (l like) Delete(ctx context.Context, userID string, jokeID string) error {
	like, err := sqlboiler.FindLike(ctx, l.executor, userID, jokeID)
	if err != nil {
		return err
	}
	_, err = like.Delete(ctx, l.executor)
	return errors.HandleErr(err)
}