package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/port/http/dto"
)

type SignHandler struct {
	userRepo    repository.UserRepository
	signUsecase application.SignUseCase
}

func NewSignHandler(
	userRepo repository.UserRepository,
	signUsecase application.SignUseCase,
) SignHandler {
	return SignHandler{
		userRepo:    userRepo,
		signUsecase: signUsecase,
	}
}

func (h *SignHandler) SignUp(ctx context.Context, req *ogen.SignUpRequest) error {
	return h.signUsecase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
}

func (h *SignHandler) SignIn(ctx context.Context, req *ogen.SignInRequest) (*ogen.SignInResponse, error) {
	user, access, refresh, err := h.signUsecase.SignIn(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &ogen.SignInResponse{
		UserResponse: dto.NewUserResponse(user),
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
