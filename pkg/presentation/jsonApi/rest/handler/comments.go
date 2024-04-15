package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
)

type CommentsHandler struct {
	comment entity.IComment
}

func NewCommentsHandler(
	comment entity.IComment,
) CommentsHandler {
	return CommentsHandler{
		comment: comment,
	}
}

func (h CommentsHandler) CommentsByJokeIDGet(ctx context.Context, params ogen.CommentsByJokeIDGetParams) (*ogen.Comments, error) {
	comments, err := h.comment.GetAllByJokeID(
		ctx,
		params.JokeID,
		restDto.NewDomainSelectParams(&params.SelectParams),
	)
	return restDto.NewCommentsResponse(comments), err
}