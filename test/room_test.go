package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestRoomCreate(t *testing.T) {
	ctx, r, entities := NewTest(t)

	user, err := entities.User.Create(ctx, entity.UserCreateRequest{
		Username: "testUsername",
		Email:    "testEmail@gmail.com",
		Password: "testPassword",
		Fullname: "testFullname",
		Status:   "testStatus",
	})
	r.NoError(err)
	r.NotEmpty(user)

	expected, err := entities.Room.Create(ctx, entity.RoomCreateRequest{
		Name:          "testName",
		Description:   "testDescription",
		CreatorUserID: user.ID,
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.Room.GetByID(ctx, expected.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(expected, actual)
}

func TestRoomUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestRoomDelete(t *testing.T) {
	t.Skip("unimplemented")
}
