package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
)

type SignHandler struct {
	authManager auth.Manager
	signUsecase application.ISignUseCase
}

func NewSignHandler(
	authManager auth.Manager,
	signUsecase application.ISignUseCase,
) SignHandler {
	return SignHandler{
		authManager: authManager,
		signUsecase: signUsecase,
	}
}

func (h SignHandler) SignUpPost(ctx context.Context, req *ogen.SignUpPostReq) error {
	_, err := h.signUsecase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	return err
}

func (h SignHandler) SignInPost(ctx context.Context, req *ogen.SignInPostReq) (*ogen.SignInPostOK, error) {
	user, err := h.signUsecase.SignIn(ctx, application.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	access, err := h.authManager.CreateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refresh, err := h.authManager.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &ogen.SignInPostOK{
		UserResponse: *restDto.NewUser(user),
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (h SignHandler) SignRefreshPost(ctx context.Context, req *ogen.SignRefreshPostReq) (*ogen.SignRefreshPostOK, error) {
	access, err := h.authManager.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &ogen.SignRefreshPostOK{
		AccessToken: access,
	}, nil
}
