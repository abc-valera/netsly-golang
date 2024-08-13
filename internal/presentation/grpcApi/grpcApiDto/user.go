package grpcApiDto

import (
	"github.com/abc-valera/netsly-golang/gen/pb"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewUserResponse(user model.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
		Status:   user.Status,
	}
}
