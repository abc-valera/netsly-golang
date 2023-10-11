package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto/common"
)

func NewUserResponse(user *entity.User) *ogen.User {
	if user == nil {
		return &ogen.User{}
	}
	return &ogen.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Fullname:  common.NewOptString(user.Fullname),
		Status:    common.NewOptString(user.Status),
		CreatedAt: user.CreatedAt,
	}
}
