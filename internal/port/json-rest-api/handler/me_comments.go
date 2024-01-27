package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domainval"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/query"
	"github.com/abc-valera/flugo-api-golang/internal/port/json-rest-api/dto"
)

type MeCommentsHandler struct {
	commentQuery  query.IComment
	commentDomain domainval.Comment
}

func NewMeCommentsHandler(
	commentQuery query.IComment,
	commentDomain domainval.Comment,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentQuery:  commentQuery,
		commentDomain: commentDomain,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) error {
	return h.commentDomain.Create(ctx, domainval.CommentCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
		Text:   req.Text,
	})
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) error {
	return h.commentDomain.Update(ctx, req.CommentID, domainval.CommentUpdateRequest{
		Text: dto.NewPointerString(req.Text),
	})
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.commentDomain.Delete(ctx, req.CommentID)
}
