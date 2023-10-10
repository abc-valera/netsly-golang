package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

func NewUserResponse(user *entity.User) *ogen.User {
	if user == nil {
		return &ogen.User{}
	}
	return &ogen.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: ogen.NewOptString(user.Fullname),
		Status:   ogen.NewOptString(user.Status),
	}
}
