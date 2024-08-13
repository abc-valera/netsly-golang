package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

var ErrRoomNameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Room already exists")

type IRoom interface {
	Create(ctx context.Context, req RoomCreateRequest) (model.Room, error)
	Update(ctx context.Context, id string, req RoomUpdateRequest) (model.Room, error)
	Delete(ctx context.Context, id string) error
}

type RoomCreateRequest struct {
	Room          model.Room
	CreatorUserID string
}

type RoomUpdateRequest struct {
	Description *string
}
