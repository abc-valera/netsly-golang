package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
)

type IComment interface {
	Create(ctx context.Context, req model.Comment) error
	Update(ctx context.Context, commentID string, req CommentUpdate) error
	Delete(ctx context.Context, id string) error
}

type CommentUpdate struct {
	Text *string
}
