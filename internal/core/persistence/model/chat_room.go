package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model/common"
)

var (
	ErrChatRoomNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatRoom not found")
)

type ChatRoom struct {
	common.BaseModel
	Name        string
	Description string
}

type ChatRooms []ChatRoom
