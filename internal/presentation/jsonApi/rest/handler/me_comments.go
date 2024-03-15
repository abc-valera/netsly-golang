package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/dto"
)

type MeCommentsHandler struct {
	commentQuery  query.IComment
	commentDomain entity.IComment
}

func NewMeCommentsHandler(
	commentQuery query.IComment,
	commentDomain entity.IComment,
) MeCommentsHandler {
	return MeCommentsHandler{
		commentQuery:  commentQuery,
		commentDomain: commentDomain,
	}
}

func (h MeCommentsHandler) MeCommentsPost(ctx context.Context, req *ogen.MeCommentsPostReq) (*ogen.Comment, error) {
	comment, err := h.commentDomain.Create(ctx, entity.CommentCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
		Text:   req.Text,
	})
	if err != nil {
		return nil, err
	}
	return dto.NewCommentResponse(comment), err
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) (*ogen.Comment, error) {
	comment, err := h.commentDomain.Update(ctx, req.CommentID, entity.CommentUpdateRequest{
		Text: dto.NewPointerString(req.Text),
	})
	return dto.NewCommentResponse(comment), err
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.commentDomain.Delete(ctx, req.CommentID)
}
