package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
)

type ILike interface {
	Create(ctx context.Context, req LikeCreateRequest) (model.Like, error)
	Delete(ctx context.Context, req DeleteLikeRequest) error
}

type like struct {
	command command.ILike
}

func NewLike(
	command command.ILike,
) ILike {
	return like{
		command: command,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l like) Create(ctx context.Context, req LikeCreateRequest) (model.Like, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Like{}, err
	}

	// Save in the data source
	return l.command.Create(ctx, model.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: time.Now(),
	})
}

type DeleteLikeRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l like) Delete(ctx context.Context, req DeleteLikeRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	// Delete from the data source
	return l.command.Delete(ctx, req.UserID, req.JokeID)
}
