package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntUser(entUser *ent.User) model.User {
	if entUser == nil {
		return model.User{}
	}
	return model.User{
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

func FromEntUserWithErrHandle(entUser *ent.User, err error) (model.User, error) {
	return FromEntUser(entUser), errhandler.HandleErr(err)
}

func FromEntUsers(entUsers []*ent.User) model.Users {
	users := make(model.Users, len(entUsers))
	for i, entUser := range entUsers {
		users[i] = FromEntUser(entUser)
	}
	return users
}

func FromEntUsersWithErrHandle(entUsers []*ent.User, err error) (model.Users, error) {
	return FromEntUsers(entUsers), errhandler.HandleErr(err)
}
