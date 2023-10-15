package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeHandler struct {
	userRepo    repository.IUserRepository
	userUsecase application.UserUseCase
}

func NewMeHandler(
	userRepo repository.IUserRepository,
	userUsecase application.UserUseCase,
) MeHandler {
	return MeHandler{
		userRepo:    userRepo,
		userUsecase: userUsecase,
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
	return h.userUsecase.UpdateUser(ctx, application.UpdateUserRequest{
		UserID:   ctx.Value(other.PayloadKey).(service.Payload).UserID,
		Username: req.Username.Value,
		Fullanme: req.Fullname.Value,
		Status:   req.Status.Value,
	})

}

func (h MeHandler) MeDel(ctx context.Context) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	return h.userRepo.Delete(ctx, userID)
}
