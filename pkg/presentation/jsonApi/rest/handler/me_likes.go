package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
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
	_, err := h.like.Create(ctx, entity.LikeCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
	return err
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	return h.like.Delete(ctx, entity.DeleteLikeRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}
