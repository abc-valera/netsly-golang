package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IComment interface {
	Create(ctx context.Context, req model.Comment) (model.Comment, error)
	Update(ctx context.Context, ids model.Comment, req CommentUpdateRequest) (model.Comment, error)
	Delete(ctx context.Context, req model.Comment) error
}

type CommentUpdateRequest struct {
	UpdatedAt time.Time

	Text *string
}
