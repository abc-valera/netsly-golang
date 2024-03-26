package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
)

type MeLikesHandler struct {
	likeQuery  query.ILike
	likeEntity entity.ILike
}

func NewMeLikesHandler(
	likeQuery query.ILike,
	likeEntity entity.ILike,
) MeLikesHandler {
	return MeLikesHandler{
		likeQuery:  likeQuery,
		likeEntity: likeEntity,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	_, err := h.likeEntity.Create(ctx, entity.LikeCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
	return err
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	return h.likeEntity.Delete(ctx, entity.DeleteLikeRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}
