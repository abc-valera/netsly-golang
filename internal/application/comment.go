package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
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

type CreateCommentRequest struct {
	UserID string
	JokeID string
	Text   string
}

func (uc CommentUseCase) CreateComment(ctx context.Context, req CreateCommentRequest) error {
	domainComment, err := entity.NewComment(req.UserID, req.JokeID, req.Text)
	if err != nil {
		return err
	}
	return uc.commentRepo.Create(ctx, domainComment)
}

func (uc CommentUseCase) GetCommentsByJoke(ctx context.Context, jokeID string, params spec.SelectParams) (entity.Comments, error) {
	if err := entity.ValidateCommentSelectParams(params); err != nil {
		return nil, err
	}
	return uc.commentRepo.GetByJokeID(ctx, jokeID, params)
}

type UpdateCommentRequest struct {
	CommentID   string
	CommentText string
	UpdaterID   string
}

func (uc CommentUseCase) UpdateComment(ctx context.Context, req UpdateCommentRequest) error {
	domainComment, err := uc.commentRepo.GetByID(ctx, req.CommentID)
	if err != nil {
		return err
	}

	if req.UpdaterID != domainComment.UserID {
		return errCommentModifyPermissionDenied
	}

	if req.CommentText != "" {
		domainComment.Text = req.CommentText
	}

	return uc.commentRepo.Update(ctx, domainComment)
}

type DeleteCommentRequest struct {
	CommentID string
	DeleterID string
}

func (uc CommentUseCase) DeleteComment(ctx context.Context, req DeleteCommentRequest) error {
	dbComment, err := uc.commentRepo.GetByID(ctx, req.CommentID)
	if err != nil {
		return err
	}

	if dbComment.UserID != req.DeleterID {
		return errCommentModifyPermissionDenied
	}

	return uc.commentRepo.Delete(ctx, req.CommentID)
}
