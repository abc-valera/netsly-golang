package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type ILike interface {
	Create(ctx context.Context, req LikeCreateRequest) (model.Like, error)
	Delete(ctx context.Context, userID, jokeID string) error

	query.ILike
}

type like struct {
	command command.ILike
	query.ILike
}

func NewLike(
	command command.ILike,
	query query.ILike,
) ILike {
	return like{
		command: command,
		ILike:   query,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (e like) Create(ctx context.Context, req LikeCreateRequest) (model.Like, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Like{}, err
	}

	return e.command.Create(ctx, model.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: time.Now(),
	})
}

func (e like) Delete(ctx context.Context, userID string, jokeID string) error {
	if err := global.Validate().Var(userID, "required,uuid"); err != nil {
		return err
	}
	if err := global.Validate().Var(jokeID, "required,uuid"); err != nil {
		return err
	}

	return e.command.Delete(ctx, userID, jokeID)
}
