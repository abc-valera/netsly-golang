package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeLikesHandler struct {
	likeRepo    repository.ILikeRepository
	likeUseCase application.LikeUseCase
}

func NewMeLikesHandler(
	likeRepo repository.ILikeRepository,
	likeUseCase application.LikeUseCase,
) MeLikesHandler {
	return MeLikesHandler{
		likeRepo:    likeRepo,
		likeUseCase: likeUseCase,
	}
}

func (h MeLikesHandler) MeLikesPost(ctx context.Context, req *ogen.MeLikesPostReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	return h.likeUseCase.CreateLike(ctx, application.CreateLikeRequest{
		UserID: userID,
		JokeID: req.JokeID,
	})
}

func (h MeLikesHandler) MeLikesDel(ctx context.Context, req *ogen.MeLikesDelReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	return h.likeRepo.Delete(ctx, userID, req.JokeID)
}
