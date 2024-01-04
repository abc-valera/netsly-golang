package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/err-handler"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntUserToUser(entUser *ent.User) *model.User {
	if entUser == nil {
		return nil
	}
	return &model.User{
		BaseModel: common.BaseModel{
			ID:        entUser.ID,
			CreatedAt: entUser.CreatedAt,
		},
		Username:       entUser.Username,
		Email:          entUser.Email,
		HashedPassword: entUser.HashedPassword,
		Fullname:       entUser.Fullname,
		Status:         entUser.Status,
	}
}

func FromEntUserToUserWithErrHandle(entUser *ent.User, err error) (*model.User, error) {
	return FromEntUserToUser(entUser), errhandler.HandleErr(err)
}

func FromEntUsersToUsers(entUsers []*ent.User) model.Users {
	users := make(model.Users, len(entUsers))
	for i, entUser := range entUsers {
		users[i] = FromEntUserToUser(entUser)
	}
	return users
}

func FromEntUsersToUsersWithErrHandle(entUsers []*ent.User, err error) (model.Users, error) {
	return FromEntUsersToUsers(entUsers), errhandler.HandleErr(err)
}
