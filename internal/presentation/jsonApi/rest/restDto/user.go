package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewUser(user model.User) *ogen.User {
	return &ogen.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Fullname:  ogen.NewOptString(user.Fullname),
		Status:    ogen.NewOptString(user.Status),
		CreatedAt: user.CreatedAt,
	}
}
