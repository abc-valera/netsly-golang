package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity/common"
)

type ChatMsg struct {
	common.BaseEntity
	ChatRoomID string
	UserID     string
	Text       string
}

func NewChatMsg(chatRoomID, userID, text string) (*ChatMsg, error) {
	if chatRoomID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid chat room ID")
	}
	if userID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}
	if text == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid text")
	}

	return &ChatMsg{
		BaseEntity: common.NewBaseEntity(),
		ChatRoomID: chatRoomID,
		UserID:     userID,
		Text:       text,
	}, nil
}

type ChatMsgs []*ChatMsg
