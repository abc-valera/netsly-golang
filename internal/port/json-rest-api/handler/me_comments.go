package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/core/domain"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/dto"
)

type MeCommentsHandler struct {
	commentQuery  query.IComment
	commentDomain domain.Comment
}

func NewMeCommentsHandler(
	commentQuery query.IComment,
	commentDomain domain.Comment,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentQuery:  commentQuery,
		commentDomain: commentDomain,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) error {
	return h.commentDomain.Create(ctx, domain.CommentCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
		Text:   req.Text,
	})
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) error {
	return h.commentDomain.Update(ctx, req.CommentID, domain.CommentUpdateRequest{
		Text: dto.NewPointerString(req.Text),
	})
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.commentDomain.Delete(ctx, req.CommentID)
}