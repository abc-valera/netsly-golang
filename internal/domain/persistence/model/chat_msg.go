package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrChatMessageNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatMessage not found")
)

type ChatMessage struct {
	common.BaseEntity
	ChatRoomID string
	UserID     string
	Text       string
}

type ChatMessages []ChatMessage
