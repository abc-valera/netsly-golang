package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

var ErrRoomMemberAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "RoomMember already exists")

type IRoomMember interface {
	Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error)
	Delete(ctx context.Context, userID, roomID string) error
}

type RoomMemberCreateRequest struct {
	model.RoomMember
	UserID string
	RoomID string
}

type RoomMemberDeleteRequest struct {
	UserID string
	RoomID string
}
