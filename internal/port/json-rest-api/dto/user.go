package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

func NewUserResponse(user model.User) *ogen.User {
	return &ogen.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Fullname:  NewOptString(user.Fullname),
		Status:    NewOptString(user.Status),
		CreatedAt: user.CreatedAt,
	}
}
