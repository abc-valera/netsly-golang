package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeCommentsHandler struct {
	commentRepo repository.ICommentRepository
}

func NewMeCommentsHandler(
	commentRepo repository.ICommentRepository,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentRepo: commentRepo,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	comment, err := entity.NewComment(
		userID,
		req.JokeID,
		req.Text,
	)
	if err != nil {
		return err
	}
	return h.commentRepo.Create(ctx, comment)
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) error {
	updateReq, err := repository.NewCommentUpdateRequest(req.Text.Value)
	if err != nil {
		return err
	}
	return h.commentRepo.Update(ctx, req.CommentID, updateReq)
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.commentRepo.Delete(ctx, req.CommentID)
}
