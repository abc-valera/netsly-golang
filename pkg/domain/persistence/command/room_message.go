package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
)

type IRoomMessage interface {
	Create(ctx context.Context, req model.RoomMessage) (model.RoomMessage, error)
	Update(ctx context.Context, id string, req RoomMessageUpdate) (model.RoomMessage, error)
	Delete(ctx context.Context, id string) error
}

type RoomMessageUpdate struct {
	Text *string
}
