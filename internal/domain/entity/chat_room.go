package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity/common"
)

type ChatRoom struct {
	common.BaseEntity
	Name        string
	Description string
}

func NewChatRoom(name, description string) (*ChatRoom, error) {
	if name == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid name")
	}
	if description == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid description")
	}

	return &ChatRoom{
		BaseEntity:  common.NewBaseEntity(),
		Name:        name,
		Description: description,
	}, nil
}

type ChatRooms []*ChatRoom
