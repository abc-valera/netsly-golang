package comments

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

func (h CommentsHandler) CommentsByJokeIDGet(ctx context.Context, params ogen.CommentsByJokeIDGetParams) (*ogen.Comments, error) {
	baseParams, err := dto.NewDomainSelectParams(&params.SelectParams)
	if err != nil {
		return nil, err
	}
	domainComments, err := h.commentRepo.GetByJokeID(ctx, params.JokeID, baseParams)
	return dto.NewCommentsResponse(domainComments), err
}
