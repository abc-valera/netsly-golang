package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

func FromEntUserToUser(entUser *ent.User) *entity.User {
	if entUser == nil {
		return nil
	}
	return &entity.User{
		ID:             entUser.ID,
		Username:       entUser.Username,
		Email:          entUser.Email,
		HashedPassword: entUser.HashedPassword,
		Fullname:       entUser.Fullname,
		Status:         entUser.Status,
		CreatedAt:      entUser.CreatedAt,
	}
}
