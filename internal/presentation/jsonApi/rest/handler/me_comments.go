package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
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
	comment, err := h.comment.Update(ctx, req.CommentID, entity.CommentUpdateRequest{
		Text: restDto.NewDomainOptionalString(req.Text),
	})
	return restDto.NewComment(comment), err
}

func (h MeCommentsHandler) MeCommentsDel(ctx context.Context, req *ogen.MeCommentsDelReq) error {
	return h.comment.Delete(ctx, req.CommentID)
}
