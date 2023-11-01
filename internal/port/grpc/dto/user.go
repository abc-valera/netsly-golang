package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/pb"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
)

func NewUserResponse(user *entity.User) *pb.UserResponse {
	if user == nil {
		return &pb.UserResponse{}
	}
	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
		Status:   user.Status,
	}
}
