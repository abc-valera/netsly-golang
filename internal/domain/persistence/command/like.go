package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

var ErrLikeAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Like already exists")

type ILike interface {
	Create(ctx context.Context, req LikeCreateRequest) (model.Like, error)
	Delete(ctx context.Context, userID, jokeID string) error
}

type LikeCreateRequest struct {
	Like   model.Like
	UserID string
	JokeID string
}
