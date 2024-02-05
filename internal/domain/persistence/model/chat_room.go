package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrChatRoomNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatRoom not found")
)

type ChatRoom struct {
	common.BaseEntity
	Name        string
	Description string
}

type ChatRooms []ChatRoom
