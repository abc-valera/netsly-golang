package entity_test

import (
	"testing"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestUserEntity(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		ctx, r, entities := NewTest(t)

		createdUser, err := entities.User.Create(ctx, entity.UserCreateRequest{
			Username: "testUsername",
			Email:    "testEmail@gmail.com",
			Password: "testPassword",
			Fullname: "testFullname",
			Status:   "testStatus",
		})
		r.NoError(err)
		r.NotEmpty(createdUser)

		returnedUser, err := entities.User.GetByID(ctx, createdUser.ID)
		r.NoError(err)
		// r.Equal(createdUser, returnedUser)

		r.Equal(
			createdUser.CreatedAt.Truncate(time.Millisecond).UTC(),
			returnedUser.CreatedAt.Truncate(time.Millisecond).UTC(),
		)

		r.Equal(time.UnixMilli(0), time.UnixMilli(0))
	})

	t.Run("Update", func(t *testing.T) {
	})

	t.Run("Delete", func(t *testing.T) {
	})
}
