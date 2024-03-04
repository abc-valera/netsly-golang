package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/rest/dto"
)

type CommentsHandler struct {
	commentQuery  query.IComment
	commentDomain entity.Comment
}

func NewCommentsHandler(
	commentQuery query.IComment,
	commentDomain entity.Comment,
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
		dto.NewDomainSelectParams(&params.SelectParams),
	)
	return dto.NewCommentsResponse(comments), err
}
