package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/contexts"
)

type MeLikesHandler struct {
	like entity.ILike
}

func NewMeLikesHandler(
	like entity.ILike,
) MeLikesHandler {
	return MeLikesHandler{
		like: like,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return err
	}

	if _, err := h.like.Create(ctx, entity.LikeCreateRequest{
		UserID: userID,
		JokeID: req.JokeID,
	}); err != nil {
		return err
	}

	return nil
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return err
	}

	return h.like.Delete(ctx, entity.DeleteLikeRequest{
		UserID: userID,
		JokeID: req.JokeID,
	})
}
