package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
)

type LikesHandler struct {
	like entity.ILike
}

func NewLikesHandler(
	like entity.ILike,
) LikesHandler {
	return LikesHandler{
		like: like,
	}
}

func (h LikesHandler) LikesByJokeIDGet(ctx context.Context, params ogen.LikesByJokeIDGetParams) (int, error) {
	return h.like.CountByJokeID(ctx, params.JokeID)
}
