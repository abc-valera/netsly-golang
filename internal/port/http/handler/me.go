package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeHandler struct {
	userRepo repository.UserRepository
}

func NewMeHandler(userRepo repository.UserRepository) MeHandler {
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
