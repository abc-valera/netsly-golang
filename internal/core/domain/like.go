package domain

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
)

var (
	ErrLikeUserIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID for like")
	ErrLikeJokeIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid joke ID for like")
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
	UserID string
	JokeID string
}

func (l Like) Create(ctx context.Context, req LikeCreateRequest) error {
	// Validation
	if req.UserID == "" {
		return ErrLikeUserIDInvalid
	}
	if req.JokeID == "" {
		return ErrLikeJokeIDInvalid
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
	UserID string
	JokeID string
}

func (l Like) Delete(ctx context.Context, req DeleteLikeRequest) error {
	// Validation
	if req.UserID == "" {
		return ErrLikeUserIDInvalid
	}
	if req.JokeID == "" {
		return ErrLikeJokeIDInvalid
	}

	// Delete from the data source
	return l.command.Delete(ctx, req.UserID, req.JokeID)
}
