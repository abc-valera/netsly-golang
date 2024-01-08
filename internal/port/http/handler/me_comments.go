package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
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
