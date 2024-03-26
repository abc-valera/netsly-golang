package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
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
