package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
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
