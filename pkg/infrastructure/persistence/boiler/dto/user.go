package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
)

func ToDomainUser(user *sqlboiler.User) model.User {
	if user == nil {
		return model.User{}
	}

	return model.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname.String,
		Status:         user.Status.String,
		CreatedAt:      user.CreatedAt,
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
