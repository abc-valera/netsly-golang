package domain

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrLikeUserIDInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid user ID for like")
	ErrLikeJokeIDInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid joke ID for like")
)

type LikeDomain struct {
	command command.ILikeCommand
}

func NewLikeDomain(
	command command.ILikeCommand,
) LikeDomain {
	return LikeDomain{
		command: command,
	}
}

type LikeCreateRequest struct {
	UserID string
	JokeID string
}

func (l LikeDomain) Create(ctx context.Context, req LikeCreateRequest) error {
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

func (l LikeDomain) Delete(ctx context.Context, req DeleteLikeRequest) error {
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
