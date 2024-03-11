package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrRoomMessageNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "RoomMessage not found")
)

type RoomMessage struct {
	common.BaseModel
	Text string

	UserID string
	RoomID string
}

type RoomMessages []RoomMessage
