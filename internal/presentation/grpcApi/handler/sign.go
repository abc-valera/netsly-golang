package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi/grpcApiDto"
)

type SignHandler struct {
	signUsecase application.ISignUsecase
	pb.UnimplementedSignServiceServer
}

func NewSignHandler(
	signUsecase application.ISignUsecase,
) pb.SignServiceServer {
	return SignHandler{
		signUsecase: signUsecase,
	}
}

func (h SignHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	_, err := h.signUsecase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	return &pb.SignUpResponse{}, handleErr(err)
}

func (h SignHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := h.signUsecase.SignIn(ctx, application.SignInRequest{
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
