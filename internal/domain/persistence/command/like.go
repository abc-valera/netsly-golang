package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

var ErrLikeAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Like already exists")

type ILike interface {
	Create(ctx context.Context, req model.Like) (model.Like, error)
	Delete(ctx context.Context, userID, jokeID string) error
}
