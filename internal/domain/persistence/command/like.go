package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

var (
	ErrLikeAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Like already exists")
)

type ILike interface {
	Create(ctx context.Context, req model.Like) (model.Like, error)
	Delete(ctx context.Context, userID, jokeID string) error
}
