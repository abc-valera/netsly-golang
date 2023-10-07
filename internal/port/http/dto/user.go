package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

func NewUserResponse(user *entity.User) ogen.UserResponse {
	if user == nil {
		return ogen.UserResponse{}
	}
	return ogen.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: ogen.NewOptString(user.Fullname),
		Status:   ogen.NewOptString(user.Status),
	}
}
