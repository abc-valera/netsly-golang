package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Like struct {
	command command.ILike

	timeMaker service.ITimeMaker
}

func NewLike(
	command command.ILike,
	timer service.ITimeMaker,
) Like {
	return Like{
		command:   command,
		timeMaker: timer,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l Like) Create(ctx context.Context, req LikeCreateRequest) (model.Like, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Like{}, err
	}

	// Save in the data source
	return l.command.Create(ctx, model.Like{
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		CreatedAt: l.timeMaker.Now(),
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
