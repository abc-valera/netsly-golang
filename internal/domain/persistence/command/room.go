package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

var (
	ErrRoomNameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Room already exists")
)

type IRoom interface {
	Create(ctx context.Context, req model.Room) (model.Room, error)
	Update(ctx context.Context, id string, req RoomUpdate) (model.Room, error)
	Delete(ctx context.Context, id string) error
}

type RoomUpdate struct {
	Description *string
}
