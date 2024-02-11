package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type Like struct {
	command command.ILike
}

func NewLike(
	command command.ILike,
) Like {
	return Like{
		command: command,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l Like) Create(ctx context.Context, req LikeCreateRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	// Domain logic
	createdAt := time.Now()

	// Save in the data source
	return l.command.Create(ctx, model.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: createdAt,
	})
}

type DeleteLikeRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l Like) Delete(ctx context.Context, req DeleteLikeRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	// Delete from the data source
	return l.command.Delete(ctx, req.UserID, req.JokeID)
}
