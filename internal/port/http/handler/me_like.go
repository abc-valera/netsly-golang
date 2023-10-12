package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type MeLikeHandler struct {
	likeRepository repository.ILikeRepository
}

func NewMeLikeHandler(
	likeRepository repository.ILikeRepository,
) MeLikeHandler {
	return MeLikeHandler{
		likeRepository: likeRepository,
	}
}

func (h MeLikeHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	domainLike, err := entity.NewLike(userID, req.JokeID)
	if err != nil {
		return err
	}
	return h.likeRepository.Create(ctx, domainLike)
}

func (h MeLikeHandler) MeLikesDelete(ctx context.Context, req *ogen.MeLikesDeleteReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	return h.likeRepository.Delete(ctx, userID, req.JokeID)
}
