package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/model/common"
)

var (
	ErrChatMessageNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatMessage not found")
)

type ChatMessage struct {
	common.BaseModel
	ChatRoomID string
	UserID     string
	Text       string
}

type ChatMessages []ChatMessage
