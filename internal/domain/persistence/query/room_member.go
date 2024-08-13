package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type IRoomMember interface {
	GetByIDs(ctx context.Context, userID, roomID string) (model.RoomMember, error)
	GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMembers, error)
}
