package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

type LikesHandler struct {
	repository.ILikeRepository
}

func NewLikesHandler(likeRepository repository.ILikeRepository) LikesHandler {
	return LikesHandler{
		ILikeRepository: likeRepository,
	}
}

func (h LikesHandler) LikesJokeIDGet(ctx context.Context, params ogen.LikesJokeIDGetParams) (int, error) {
	return h.ILikeRepository.CountByJokeID(ctx, params.JokeID)
}
