package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeHandler struct {
	userRepo    repository.IUserRepository
	userUseCase application.UserUseCase
}

func NewMeHandler(
	userRepo repository.IUserRepository,
	userUseCase application.UserUseCase,
) MeHandler {
	return MeHandler{
		userRepo: userRepo,
	}
}

func (h MeHandler) MeGet(ctx context.Context) (*ogen.User, error) {
	payload := ctx.Value(PayloadKey).(service.Payload)
	domainUser, err := h.userRepo.GetByID(ctx, payload.UserID)
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(domainUser), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) error {
	payload := ctx.Value(PayloadKey).(service.Payload)
	return h.userUseCase.UpdateUser(ctx, application.UpdateUserRequest{
		UpdaterID: payload.UserID,
		UserID:    req.UserID,
		Username:  req.Username.Value,
		Fullanme:  req.Fullname.Value,
		Status:    req.Status.Value,
	})
}

func (h MeHandler) MeDelete(ctx context.Context, req *ogen.MeDeleteReq) error {
	payload := ctx.Value(PayloadKey).(service.Payload)
	return h.userUseCase.DeleteUser(ctx, application.DeleteUserRequest{
		DeleterID: payload.UserID,
		UserID:    req.UserID,
	})
}
