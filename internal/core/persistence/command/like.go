package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

var (
	ErrLikeAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Like already exists")
)

type ILike interface {
	Create(ctx context.Context, req model.Like) error
	Delete(ctx context.Context, userID, jokeID string) error
}
