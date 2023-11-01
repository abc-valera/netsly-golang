package likes

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
)

type LikesHandler struct {
	likeRepo repository.ILikeRepository
}

func NewLikesHandler(likeRepo repository.ILikeRepository) LikesHandler {
	return LikesHandler{
		likeRepo: likeRepo,
	}
}

func (h LikesHandler) LikesByJokeIDGet(ctx context.Context, params ogen.LikesByJokeIDGetParams) (int, error) {
	return h.likeRepo.CountByJokeID(ctx, params.JokeID)
}
