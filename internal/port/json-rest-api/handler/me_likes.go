package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domainval"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/query"
)

type MeLikesHandler struct {
	likeQuery  query.ILike
	likeDomain domainval.Like
}

func NewMeLikesHandler(
	likeQuery query.ILike,
	likeDomain domainval.Like,
) MeLikesHandler {
	return MeLikesHandler{
		likeQuery:  likeQuery,
		likeDomain: likeDomain,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	return h.likeDomain.Create(ctx, domainval.LikeCreateRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	return h.likeDomain.Delete(ctx, domainval.DeleteLikeRequest{
		UserID: payloadUserID(ctx),
		JokeID: req.JokeID,
	})
}
