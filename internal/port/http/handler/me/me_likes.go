package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeLikesHandler struct {
	likeRepo repository.ILikeRepository
}

func NewMeLikesHandler(
	likeRepo repository.ILikeRepository,
) MeLikesHandler {
	return MeLikesHandler{
		likeRepo: likeRepo,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	like, err := entity.NewLike(userID, req.JokeID)
	if err != nil {
		return err
	}
	return h.likeRepo.Create(ctx, like)
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	return h.likeRepo.Delete(ctx, userID, req.JokeID)
}
