package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi/grpcApiDto"
)

type SignHandler struct {
	signUseCase application.ISignUseCase
	pb.UnimplementedSignServiceServer
}

func NewSignHandler(
	signUseCase application.ISignUseCase,
) pb.SignServiceServer {
	return SignHandler{
		signUseCase: signUseCase,
	}
}

func (h SignHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	_, err := h.signUseCase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	return &pb.SignUpResponse{}, handleErr(err)
}

func (h SignHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := h.signUseCase.SignIn(ctx, application.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, handleErr(err)
	}
	return &pb.SignInResponse{
		UserResponse: grpcApiDto.NewUserResponse(user),
	}, nil
}
