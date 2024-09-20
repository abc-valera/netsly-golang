package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestUserCreate(t *testing.T) {
	ctx, r, entities := NewTest(t)

	expected, err := entities.User.Create(ctx, entity.UserCreateRequest{
		Username: "testUsername",
		Email:    "testEmail@gmail.com",
		Password: "testPassword",
		Fullname: "testFullname",
		Status:   "testStatus",
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.User.GetByID(ctx, expected.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(expected, actual)
}

func TestUserUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestUserDelete(t *testing.T) {
	t.Skip("unimplemented")
}
