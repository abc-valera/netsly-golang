package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrChatRoomNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatRoom not found")

	ErrChatRoomNameAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "ChatRoom already exists")
)

type ChatRoom struct {
	common.BaseModel
	Name        string
	Description string
}

type ChatRooms []ChatRoom
