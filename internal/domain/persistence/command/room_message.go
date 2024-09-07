package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IRoomMessage interface {
	Create(ctx context.Context, req model.RoomMessage) (model.RoomMessage, error)
	Update(ctx context.Context, ids model.RoomMessage, req RoomMessageUpdateRequest) (model.RoomMessage, error)
	Delete(ctx context.Context, req model.RoomMessage) error
}

type RoomMessageUpdateRequest struct {
	UpdatedAt time.Time

	Text *string
}
