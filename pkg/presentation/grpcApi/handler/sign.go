package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/grpcApi/grpcApiDto"
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
	err := h.signUseCase.SignUp(ctx, application.SignUpRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	return &pb.SignUpResponse{}, handleErr(err)
}

func (h SignHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	resp, err := h.signUseCase.SignIn(ctx, application.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, handleErr(err)
	}
	return &pb.SignInResponse{
		UserResponse: grpcApiDto.NewUserResponse(resp.User),
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
