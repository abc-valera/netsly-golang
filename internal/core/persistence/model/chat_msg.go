package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model/common"
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
