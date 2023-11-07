package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeHandler struct {
	userRepo repository.IUserRepository
}

func NewMeHandler(
	userRepo repository.IUserRepository,
) MeHandler {
	return MeHandler{
		userRepo: userRepo,
	}
}

func (h MeHandler) MeGet(ctx context.Context) (*ogen.User, error) {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	user, err := h.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) error {
	updateReq, err := repository.NewUserUpdateRequest(
		req.Username.Value,
		"",
		"",
		req.Fullname.Value,
		req.Status.Value,
	)
	if err != nil {
		return err
	}
	return h.userRepo.Update(
		ctx,
		ctx.Value(other.PayloadKey).(service.Payload).UserID,
		updateReq,
	)
}

func (h MeHandler) MeDel(ctx context.Context) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	return h.userRepo.Delete(ctx, userID)
}
