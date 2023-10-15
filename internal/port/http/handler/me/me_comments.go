package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeCommentsHandler struct {
	commentRepo    repository.ICommentRepository
	commentUseCase application.CommentUseCase
}

func NewMeCommentsHandler(
	commentRepo repository.ICommentRepository,
	commentUseCase application.CommentUseCase,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentRepo:    commentRepo,
		commentUseCase: commentUseCase,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) error {
	return h.commentUseCase.CreateComment(ctx, application.CreateCommentRequest{
		UserID: ctx.Value(other.PayloadKey).(service.Payload).UserID,
		JokeID: req.JokeID,
		Text:   req.Text,
	})
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) error {
	return h.commentUseCase.UpdateComment(ctx, application.UpdateCommentRequest{
		CommentID:   req.CommentID,
		CommentText: req.Text.Value,
		UpdaterID:   ctx.Value(other.PayloadKey).(service.Payload).UserID,
	})
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.commentUseCase.DeleteComment(ctx, application.DeleteCommentRequest{
		CommentID: req.CommentID,
		DeleterID: ctx.Value(other.PayloadKey).(service.Payload).UserID,
	})
}
