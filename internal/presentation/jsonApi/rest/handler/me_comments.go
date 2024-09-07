package handler

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restDto"
	"go.opentelemetry.io/otel/trace"
)

type MeCommentsHandler struct {
	comment entity.IComment
}

func NewMeCommentsHandler(
	comment entity.IComment,
) MeCommentsHandler {
	return MeCommentsHandler{
		comment: comment,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) (*ogen.Comment, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	comment, err := h.comment.Create(ctx, entity.CommentCreateRequest{
		UserID: userID,
		JokeID: req.JokeID,
		Text:   req.Text,
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewComment(comment), err
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) (*ogen.Comment, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	comment, err := h.comment.Update(ctx, req.CommentID, entity.CommentUpdateRequest{
		Text: restDto.NewDomainOptionalString(req.Text),
	})
	return restDto.NewComment(comment), err
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return h.comment.Delete(ctx, req.CommentID)
}
