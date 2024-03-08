package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type Comment struct {
	command command.IComment
}

func NewComment(
	command command.IComment,
) Comment {
	return Comment{
		command: command,
	}
}

type CommentCreateRequest struct {
	Text   string `validate:"required,min=4,max=256"`
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (c Comment) Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Comment{}, err
	}

	baseModel := common.NewBaseEntity()

	return c.command.Create(ctx, model.Comment{
		BaseEntity: baseModel,
		Text:       req.Text,
		UserID:     req.UserID,
		JokeID:     req.JokeID,
	})
}

type CommentUpdateRequest struct {
	Text *string `validate:"min=4,max=256"`
}

func (c Comment) Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Comment{}, err
	}

	return c.command.Update(ctx, commentID, command.CommentUpdate{
		Text: req.Text,
	})
}

func (c Comment) Delete(ctx context.Context, commentID string) error {
	if err := global.Validator().Var(commentID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, commentID)
}
