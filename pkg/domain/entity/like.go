package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
)

type ILike interface {
	Create(ctx context.Context, req LikeCreateRequest) (model.Like, error)
	Delete(ctx context.Context, req DeleteLikeRequest) error

	query.ILike
}

type like struct {
	command command.ILike
	query.ILike

	validator service.IValidator
}

func NewLike(
	command command.ILike,
	query query.ILike,
	validator service.IValidator,
) ILike {
	return like{
		command:   command,
		ILike:     query,
		validator: validator,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (l like) Create(ctx context.Context, req LikeCreateRequest) (model.Like, error) {
	if err := l.validator.Struct(req); err != nil {
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
	if err := l.validator.Struct(req); err != nil {
		return err
	}

	// Delete from the data source
	return l.command.Delete(ctx, req.UserID, req.JokeID)
}
