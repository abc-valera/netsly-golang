package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type CommentsHandler struct {
	commentRepo repository.ICommentRepository
}

func NewCommentsHandler(commentRepo repository.ICommentRepository) CommentsHandler {
	return CommentsHandler{
		commentRepo: commentRepo,
	}
}

func (h CommentsHandler) CommentsJokeIDGet(ctx context.Context, params ogen.CommentsJokeIDGetParams) (*ogen.Comments, error) {
	domainComments, err := h.commentRepo.GetByJokeID(ctx, params.JokeID)
	return dto.NewCommentsResponse(domainComments), err
}
