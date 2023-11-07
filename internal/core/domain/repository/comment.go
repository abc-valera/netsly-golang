package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/spec"
)

var (
	ErrCommentNotFound = codeerr.NewMsgErr(codeerr.CodeNotFound, "Comment not found")

	ErrCommentsOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'created_at' field")
)

type ICommentRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByJokeID(ctx context.Context, jokeID string, spec spec.SelectParams) (entity.Comments, error)
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, commentID string, req CommentUpdateRequest) error
	Delete(ctx context.Context, commentID string) error

	common.Transactioneer
}

func ValidateCommentSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "created_at" {
		return ErrCommentsOrderByNotSupported
	}
	return nil
}

type CommentUpdateRequest struct {
	Text string
}

func NewCommentUpdateRequest(text string) (CommentUpdateRequest, error) {
	if text == "" {
		return CommentUpdateRequest{}, entity.ErrCommentTextInvalid
	}

	return CommentUpdateRequest{
		Text: text,
	}, nil
}
