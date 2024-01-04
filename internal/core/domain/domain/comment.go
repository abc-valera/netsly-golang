package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrCommentIDInvalid     = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid comment ID")
	ErrCommentUserIDInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid user ID for comment")
	ErrCommentJokeIDInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid joke ID for comment")
	ErrCommentTextInvalid   = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid text")
)

type CommentDomain struct {
	command command.ICommentCommand
}

func NewCommentDomain(
	command command.ICommentCommand,
) CommentDomain {
	return CommentDomain{
		command: command,
	}
}

type CommentCreateRequest struct {
	UserID string
	JokeID string
	Text   string
}

func (c CommentDomain) Create(ctx context.Context, req CommentCreateRequest) error {
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

func (c CommentDomain) Update(ctx context.Context, commentID string, req CommentUpdateRequest) error {
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

func (c CommentDomain) Delete(ctx context.Context, commentID string) error {
	// Validation
	if commentID == "" {
		return ErrCommentIDInvalid
	}

	return c.command.Delete(ctx, commentID)
}