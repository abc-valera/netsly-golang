package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/model"
)

type ILike interface {
	Create(ctx context.Context, req model.Like) error
	Delete(ctx context.Context, userID, jokeID string) error
}
