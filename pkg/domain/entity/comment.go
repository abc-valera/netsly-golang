package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/google/uuid"
)

type IComment interface {
	Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error)
	Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error)
	Delete(ctx context.Context, commentID string) error
}

type comment struct {
	command command.IComment
}

func NewComment(
	command command.IComment,
) IComment {
	return comment{
		command: command,
	}
}

type CommentCreateRequest struct {
	Text   string `validate:"required,min=4,max=256"`
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (c comment) Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Comment{}, err
	}

	return c.command.Create(ctx, model.Comment{
		ID:        uuid.New().String(),
		Text:      req.Text,
		CreatedAt: time.Now(),
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	})
}

type CommentUpdateRequest struct {
	Text *string `validate:"min=4,max=256"`
}

func (c comment) Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Comment{}, err
	}

	return c.command.Update(ctx, commentID, command.CommentUpdate{
		Text: req.Text,
	})
}

func (c comment) Delete(ctx context.Context, commentID string) error {
	if err := global.Validator().Var(commentID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, commentID)
}
