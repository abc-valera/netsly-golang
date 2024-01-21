package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
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
