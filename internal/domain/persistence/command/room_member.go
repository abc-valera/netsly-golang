package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrRoomMemberAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "RoomMember already exists")

type IRoomMember interface {
	Create(ctx context.Context, req model.RoomMember) (model.RoomMember, error)
	Delete(ctx context.Context, req model.RoomMember) error
}
