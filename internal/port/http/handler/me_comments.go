// Write the code for the MeCommentsHandler as in the file internal/port/http/handler/me_jokes.go

// Path: internal/port/http/handler/me_jokes.go
// Compare this snippet from internal/port/http/handler/me_jokes.go:
package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type MeCommentsHandler struct {
	commentRepo    repository.ICommentRepository
	commentUsecase application.CommentUseCase
}

func NewMeCommentsHandler(
	commentRepo repository.ICommentRepository,
	commentUsecase application.CommentUseCase,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentRepo:    commentRepo,
		commentUsecase: commentUsecase,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	domainJoke, err := entity.NewComment(userID, req.JokeID, req.Text)
	if err != nil {
		return err
	}
	return h.commentRepo.Create(ctx, domainJoke)
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	return h.commentUsecase.UpdateComment(ctx, application.UpdateCommentRequest{
		CommentID:   req.CommentID,
		CommentText: req.Text.Value,
		UserID:      userID,
	})
}

func (h MeCommentsHandler) MeCommentsDelete(ctx context.Context, req *ogen.MeCommentsDeleteReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	return h.commentUsecase.DeleteComment(ctx, application.DeleteCommentRequest{
		CommentID: req.CommentID,
		UserID:    userID,
	})
}
