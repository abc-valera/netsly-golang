package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
)

func ToDomainUser(user *sqlboiler.User) model.User {
	if user == nil {
		return model.User{}
	}

	return model.User{
		BaseModel: common.BaseModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
		},
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname.String,
		Status:         user.Status.String,
	}
}

func ToDomainUserWithErrHandle(user *sqlboiler.User, err error) (model.User, error) {
	return ToDomainUser(user), errors.HandleErr(err)
}

func ToDomainUsers(users sqlboiler.UserSlice) model.Users {
	var domainUsers model.Users
	for _, user := range users {
		domainUsers = append(domainUsers, ToDomainUser(user))
	}
	return domainUsers
}

func ToDomainUsersWithErrHandle(users sqlboiler.UserSlice, err error) (model.Users, error) {
	return ToDomainUsers(users), errors.HandleErr(err)
}
