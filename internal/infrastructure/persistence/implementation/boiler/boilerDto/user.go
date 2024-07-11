package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
)

func NewDomainUser(user *sqlboiler.User) model.User {
	if user == nil {
		return model.User{}
	}

	return model.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt,
	}
}

func NewDomainUserWithErrHandle(user *sqlboiler.User, err error) (model.User, error) {
	return NewDomainUser(user), errutil.HandleErr(err)
}

func NewDomainUsers(users sqlboiler.UserSlice) model.Users {
	var domainUsers model.Users
	for _, user := range users {
		domainUsers = append(domainUsers, NewDomainUser(user))
	}
	return domainUsers
}

func NewDomainUsersWithErrHandle(users sqlboiler.UserSlice, err error) (model.Users, error) {
	return NewDomainUsers(users), errutil.HandleErr(err)
}
