package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrRoomNameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Room already exists")

type IRoom interface {
	Create(ctx context.Context, req model.Room) (model.Room, error)
	Update(ctx context.Context, ids model.Room, req RoomUpdateRequest) (model.Room, error)
	Delete(ctx context.Context, req model.Room) error
}

type RoomUpdateRequest struct {
	UpdatedAt time.Time

	Description *string
}
