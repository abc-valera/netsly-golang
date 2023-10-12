package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

var (
	errCommentModifyPermissionDenied = codeerr.NewMsgErr(codeerr.CodePermissionDenied, "You can modify only your own comments")
)

type CommentUseCase struct {
	commentRepo repository.ICommentRepository
}

func NewCommentUseCase(commentRepo repository.ICommentRepository) CommentUseCase {
	return CommentUseCase{
		commentRepo: commentRepo,
	}
}

type UpdateCommentRequest struct {
	CommentID   string
	CommentText string
	UserID      string
}

func (uc CommentUseCase) UpdateComment(ctx context.Context, req UpdateCommentRequest) error {
	domainComment, err := uc.commentRepo.GetByID(ctx, req.CommentID)
	if err != nil {
		return err
	}

	if req.UserID != domainComment.UserID {
		return errCommentModifyPermissionDenied
	}

	if req.CommentText != "" {
		domainComment.Text = req.CommentText
	}

	return uc.commentRepo.Update(ctx, domainComment)
}

type DeleteCommentRequest struct {
	CommentID string
	UserID    string
}

func (uc CommentUseCase) DeleteComment(ctx context.Context, req DeleteCommentRequest) error {
	dbComment, err := uc.commentRepo.GetByID(ctx, req.CommentID)
	if err != nil {
		return err
	}

	if dbComment.UserID != req.UserID {
		return errCommentModifyPermissionDenied
	}

	return uc.commentRepo.Delete(ctx, req.CommentID)
}
