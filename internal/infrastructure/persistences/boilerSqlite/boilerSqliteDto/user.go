package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
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
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

func NewDomainUsers(users sqlboiler.UserSlice) model.Users {
	var domainUsers model.Users
	for _, user := range users {
		domainUsers = append(domainUsers, NewDomainUser(user))
	}
	return domainUsers
}
