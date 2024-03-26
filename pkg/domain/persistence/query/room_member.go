package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
)

type IRoomMember interface {
	GetByIDs(ctx context.Context, userID, roomID string) (model.RoomMember, error)
	GetByRoomID(ctx context.Context, roomID string) (model.RoomMembers, error)
}
