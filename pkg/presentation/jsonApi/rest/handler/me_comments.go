package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
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
	comment, err := h.comment.Create(ctx, entity.CommentCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
		Text:   req.Text,
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewCommentResponse(comment), err
}

func (h MeCommentsHandler) MeCommentsPut(ctx context.Context, req *ogen.MeCommentsPutReq) (*ogen.Comment, error) {
	comment, err := h.comment.Update(ctx, req.CommentID, entity.CommentUpdateRequest{
		Text: restDto.NewPointerString(req.Text),
	})
	return restDto.NewCommentResponse(comment), err
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.comment.Delete(ctx, req.CommentID)
}