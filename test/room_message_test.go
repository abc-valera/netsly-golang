package test

import (
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
)

func TestRoomMessageCreate(t *testing.T) {
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

	room, err := entities.Room.Create(ctx, entity.RoomCreateRequest{
		Name:          "testName",
		Description:   "testDescription",
		CreatorUserID: user.ID,
	})
	r.NoError(err)
	r.NotEmpty(room)

	expected, err := entities.RoomMessage.Create(ctx, entity.RoomMessageCreateRequest{
		Text:   "testText",
		UserID: user.ID,
		RoomID: room.ID,
	})
	r.NoError(err)
	r.NotEmpty(expected)

	actual, err := entities.RoomMessage.GetByID(ctx, expected.ID)
	r.NoError(err)
	r.NotEmpty(actual)

	r.Equal(expected, actual)
}

func TestRoomMessageUpdate(t *testing.T) {
	t.Skip("unimplemented")
}

func TestRoomMessageDelete(t *testing.T) {
	t.Skip("unimplemented")
}
