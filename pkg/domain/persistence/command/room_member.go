package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
)

var (
	ErrRoomMemberAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "RoomMember already exists")
)

type IRoomMember interface {
	Create(ctx context.Context, req model.RoomMember) (model.RoomMember, error)
	Delete(ctx context.Context, RoomID, UserID string) error
}
