package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrChatRoomNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "ChatRoom not found")

	ErrChatRoomNameAlreadyExists = codeerr.NewMessageErr(codeerr.CodeAlreadyExists, "ChatRoom already exists")
)

type ChatRoom struct {
	common.BaseModel
	Name        string
	Description string
}

type ChatRoomUpdate struct {
	Description *string
}

type ChatRooms []*ChatRoom
