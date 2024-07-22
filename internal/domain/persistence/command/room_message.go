package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type IRoomMessage interface {
	Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error)
	Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error)
	Delete(ctx context.Context, id string) error
}

type RoomMessageCreateRequest struct {
	RoomMessage model.RoomMessage
	UserID      string
	RoomID      string
}

type RoomMessageUpdateRequest struct {
	Text *string
}
