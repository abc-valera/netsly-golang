package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type MeLikesHandler struct {
	likeQuery  query.ILike
	likeDomain entity.Like
}

func NewMeLikesHandler(
	likeQuery query.ILike,
	likeDomain entity.Like,
) MeLikesHandler {
	return MeLikesHandler{
		likeQuery:  likeQuery,
		likeDomain: likeDomain,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	return h.likeDomain.Create(ctx, entity.LikeCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	return h.likeDomain.Delete(ctx, entity.DeleteLikeRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}
