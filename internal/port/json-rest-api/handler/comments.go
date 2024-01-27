package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query"
	"github.com/abc-valera/flugo-api-golang/internal/port/json-rest-api/dto"
)

type CommentsHandler struct {
	commentQuery  query.IComment
	commentDomain domain.Comment
}

func NewCommentsHandler(
	commentQuery query.IComment,
	commentDomain domain.Comment,
) CommentsHandler {
	return CommentsHandler{
		commentQuery:  commentQuery,
		commentDomain: commentDomain,
	}
}

func (h CommentsHandler) CommentsByJokeIDGet(ctx context.Context, params ogen.CommentsByJokeIDGetParams) (*ogen.Comments, error) {
	slectParams, err := dto.NewDomainSelectParams(&params.SelectParams)
	if err != nil {
		return nil, err
	}
	comments, err := h.commentQuery.GetAllByJokeID(ctx, params.JokeID, slectParams)
	return dto.NewCommentsResponse(comments), err
}
