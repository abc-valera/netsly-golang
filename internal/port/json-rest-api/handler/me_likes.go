package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
)

type MeLikesHandler struct {
	likeQuery  query.ILike
	likeDomain domain.Like
}

func NewMeLikesHandler(
	likeQuery query.ILike,
	likeDomain domain.Like,
) MeLikesHandler {
	return MeLikesHandler{
		likeQuery:  likeQuery,
		likeDomain: likeDomain,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	return h.likeDomain.Create(ctx, domain.LikeCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	return h.likeDomain.Delete(ctx, domain.DeleteLikeRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}
