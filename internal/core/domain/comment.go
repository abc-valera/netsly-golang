package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
)

var (
	ErrCommentIDInvalid     = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid comment ID")
	ErrCommentUserIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID for comment")
	ErrCommentJokeIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid joke ID for comment")
	ErrCommentTextInvalid   = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid text")
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
	UserID string
	JokeID string
	Text   string
}

func (c Comment) Create(ctx context.Context, req CommentCreateRequest) error {
	// Validation
	if req.UserID == "" {
		return ErrCommentUserIDInvalid
	}
	if req.JokeID == "" {
		return ErrCommentJokeIDInvalid
	}
	if req.Text == "" || len(req.Text) < 4 || len(req.Text) > 256 {
		return ErrCommentTextInvalid
	}

	// Domain logic
	baseModel := common.NewBaseModel()

	return c.command.Create(ctx, model.Comment{
		BaseModel: baseModel,
		UserID:    req.UserID,
		JokeID:    req.JokeID,
		Text:      req.Text,
	})
}

type CommentUpdateRequest struct {
	Text *string
}

func (c Comment) Update(ctx context.Context, commentID string, req CommentUpdateRequest) error {
	// Validation
	if commentID == "" {
		return ErrCommentIDInvalid
	}
	if req.Text != nil && (len(*req.Text) < 4 || len(*req.Text) > 256) {
		return ErrCommentTextInvalid
	}

	// Domain logic
	return c.command.Update(ctx, commentID, command.CommentUpdate{
		Text: req.Text,
	})
}

func (c Comment) Delete(ctx context.Context, commentID string) error {
	// Validation
	if commentID == "" {
		return ErrCommentIDInvalid
	}

	return c.command.Delete(ctx, commentID)
}
