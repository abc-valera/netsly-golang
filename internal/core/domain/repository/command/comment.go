package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type ICommentCommand interface {
	Create(ctx context.Context, req model.Comment) error
	Update(ctx context.Context, commentID string, req model.CommentUpdate) error
	Delete(ctx context.Context, id string) error
}
