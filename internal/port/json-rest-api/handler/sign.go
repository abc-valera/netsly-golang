package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/dto"
)

type SignHandler struct {
	signUsecase application.SignUseCase
}

func NewSignHandler(
	signUsecase application.SignUseCase,
) SignHandler {
	return SignHandler{
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
