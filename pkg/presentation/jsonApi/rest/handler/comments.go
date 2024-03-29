package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
)

type CommentsHandler struct {
	commentQuery  query.IComment
	commentDomain entity.IComment
}

func NewCommentsHandler(
	commentQuery query.IComment,
	commentDomain entity.IComment,
) CommentsHandler {
	return CommentsHandler{
		commentQuery:  commentQuery,
		commentDomain: commentDomain,
	}
}

func (h CommentsHandler) CommentsByJokeIDGet(ctx context.Context, params ogen.CommentsByJokeIDGetParams) (*ogen.Comments, error) {
	comments, err := h.commentQuery.GetAllByJokeID(
		ctx,
		params.JokeID,
		restDto.NewDomainSelectParams(&params.SelectParams),
	)
	return restDto.NewCommentsResponse(comments), err
}
