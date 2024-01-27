package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/query"
)

type LikesHandler struct {
	likeQuery query.ILike
}

func NewLikesHandler(
	likeRepo query.ILike,
) LikesHandler {
	return LikesHandler{
		likeQuery: likeRepo,
	}
}

func (h LikesHandler) LikesByJokeIDGet(ctx context.Context, params ogen.LikesByJokeIDGetParams) (int, error) {
	return h.likeQuery.CountByJokeID(ctx, params.JokeID)
}
