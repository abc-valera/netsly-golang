package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/core/optional"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
)

type IComment interface {
	Create(ctx context.Context, req model.Comment) (model.Comment, error)
	Update(ctx context.Context, commentID string, req CommentUpdate) (model.Comment, error)
	Delete(ctx context.Context, id string) error
}

type CommentUpdate struct {
	Text optional.Optional[string]
}
