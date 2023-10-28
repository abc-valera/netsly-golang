package sign

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type SignHandler struct {
	userRepo    repository.IUserRepository
	signUsecase application.SignUseCase
}

func NewSignHandler(
	userRepo repository.IUserRepository,
	signUsecase application.SignUseCase,
) SignHandler {
	return SignHandler{
		userRepo:    userRepo,
		signUsecase: signUsecase,
	}
}

func (h SignHandler) SignUpPost(ctx context.Context, req *ogen.SignUpPostReq) error {
	return h.signUsecase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
}

func (h SignHandler) SignInPost(ctx context.Context, req *ogen.SignInPostReq) (*ogen.SignInPostOK, error) {
	resp, err := h.signUsecase.SignIn(ctx, application.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &ogen.SignInPostOK{
		UserResponse: *dto.NewUserResponse(resp.User),
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (h SignHandler) SignRefreshPost(ctx context.Context, req *ogen.SignRefreshPostReq) (*ogen.SignRefreshPostOK, error) {
	access, err := h.signUsecase.SignRefresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &ogen.SignRefreshPostOK{
		AccessToken: access,
	}, nil
}