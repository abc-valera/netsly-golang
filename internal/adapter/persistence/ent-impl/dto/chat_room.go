package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

func FromEntChatRoom(entChatRoom *ent.ChatRoom) model.ChatRoom {
	if entChatRoom == nil {
		return model.ChatRoom{}
	}
	return model.ChatRoom{
		BaseEntity: common.BaseEntity{
			ID:        entChatRoom.ID,
			CreatedAt: entChatRoom.CreatedAt,
		},
		Name:        entChatRoom.Name,
		Description: entChatRoom.Description,
	}
}

func FromEntChatRoomWithErrHandle(entChatRoom *ent.ChatRoom, err error) (model.ChatRoom, error) {
	return FromEntChatRoom(entChatRoom), err
}

func FromEntChatRooms(entChatRooms []*ent.ChatRoom) model.ChatRooms {
	chatRooms := make(model.ChatRooms, len(entChatRooms))
	for i, entChatRoom := range entChatRooms {
		chatRooms[i] = FromEntChatRoom(entChatRoom)
	}
	return chatRooms
}

func FromEntChatRoomsWithErrHandle(entChatRooms []*ent.ChatRoom, err error) (model.ChatRooms, error) {
	return FromEntChatRooms(entChatRooms), err
}
