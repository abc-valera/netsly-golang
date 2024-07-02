package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
)

type IComment interface {
	Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error)
	Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error)
	Delete(ctx context.Context, commentID string) error

	query.IComment
}

type comment struct {
	command command.IComment
	query.IComment
}

func NewComment(
	command command.IComment,
	query query.IComment,
) IComment {
	return comment{
		command:  command,
		IComment: query,
	}
}

type CommentCreateRequest struct {
	Text string `validate:"required,min=4,max=256"`

	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (e comment) Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Comment{}, err
	}

	return e.command.Create(ctx, req.UserID, req.JokeID, model.Comment{
		ID:        uuid.New().String(),
		Text:      req.Text,
		CreatedAt: time.Now(),
	})
}

type CommentUpdateRequest struct {
	Text *string `validate:"min=4,max=256"`
}

func (e comment) Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Comment{}, err
	}

	return e.command.Update(ctx, commentID, command.CommentUpdate{
		Text: req.Text,
	})
}

func (e comment) Delete(ctx context.Context, commentID string) error {
	if err := global.Validate().Var(commentID, "uuid"); err != nil {
		return err
	}

	return e.command.Delete(ctx, commentID)
}
