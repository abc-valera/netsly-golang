package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntChatRoom(entChatRoom *ent.ChatRoom) model.ChatRoom {
	if entChatRoom == nil {
		return model.ChatRoom{}
	}
	return model.ChatRoom{
		BaseModel: common.BaseModel{
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
