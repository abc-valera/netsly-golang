package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type IRoomMember interface {
	GetByIDs(ctx context.Context, userID, roomID string) (model.RoomMember, error)
	GetAllByRoomID(ctx context.Context, roomID string) (model.RoomMembers, error)
}
