package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
)

func FromEntUserToUser(entUser *ent.User) *entity.User {
	if entUser == nil {
		return nil
	}
	return &entity.User{
		BaseEntity:     common.FromEntToBaseEntity(entUser.ID, entUser.CreatedAt),
		Username:       entUser.Username,
		Email:          entUser.Email,
		HashedPassword: entUser.HashedPassword,
		Fullname:       entUser.Fullname,
		Status:         entUser.Status,
	}
}

func FromEntUsersToUsers(entUsers []*ent.User) entity.Users {
	users := make(entity.Users, len(entUsers))
	for i, entUser := range entUsers {
		users[i] = FromEntUserToUser(entUser)
	}
	return users
}
