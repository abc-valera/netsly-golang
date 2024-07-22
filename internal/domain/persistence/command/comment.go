package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type IComment interface {
	Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error)
	Update(ctx context.Context, id string, req CommentUpdateRequest) (model.Comment, error)
	Delete(ctx context.Context, id string) error
}

type CommentCreateRequest struct {
	Comment model.Comment
	UserID  string
	JokeID  string
}

type CommentUpdateRequest struct {
	Text *string
}
