package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/pb"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/port/grpc/dto"
)

type SignHandler struct {
	userRepo    repository.IUserRepository
	signUsecase application.SignUseCase
	pb.UnimplementedSignServiceServer
}

func NewSignHandler(
	userRepo repository.IUserRepository,
	signUsecase application.SignUseCase,
) pb.SignServiceServer {
	return SignHandler{
		userRepo:    userRepo,
		signUsecase: signUsecase,
	}
}

func (h SignHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	err := h.signUsecase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	return &pb.SignUpResponse{}, handleErr(err)
}

func (h SignHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	resp, err := h.signUsecase.SignIn(ctx, application.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, handleErr(err)
	}
	return &pb.SignInResponse{
		UserResponse: dto.NewUserResponse(resp.User),
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
