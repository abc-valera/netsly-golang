package handler

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restDto"
	"go.opentelemetry.io/otel/trace"
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
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	comments, err := h.comment.GetAllByJokeID(
		ctx,
		params.JokeID,
		restDto.NewDomainSelector(&params.Selector),
	)
	return restDto.NewComments(comments), err
}
