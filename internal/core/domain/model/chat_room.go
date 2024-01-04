package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrChatRoomNotFound = codeerr.NewMessage(codeerr.CodeNotFound, "ChatRoom not found")

	ErrChatRoomNameAlreadyExists = codeerr.NewMessage(codeerr.CodeAlreadyExists, "ChatRoom already exists")
)

type ChatRoom struct {
	common.BaseModel
	Name        string
	Description string
}

type ChatRooms []*ChatRoom
