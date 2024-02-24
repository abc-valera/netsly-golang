package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

func FromEntRoomMessage(entRoomMessage *ent.RoomMessage) model.RoomMessage {
	if entRoomMessage == nil {
		return model.RoomMessage{}
	}
	return model.RoomMessage{
		BaseEntity: common.BaseEntity{
			ID:        entRoomMessage.ID,
			CreatedAt: entRoomMessage.CreatedAt,
		},
		UserID: entRoomMessage.Edges.User.ID,
		RoomID: entRoomMessage.Edges.Room.ID,
		Text:   entRoomMessage.Text,
	}
}

func FromEntRoomMessageWithErrHandle(entRoomMessage *ent.RoomMessage, err error) (model.RoomMessage, error) {
	return FromEntRoomMessage(entRoomMessage), err
}

func FromEntRoomMessages(entRoomMessages []*ent.RoomMessage) model.RoomMessages {
	roomMessages := make(model.RoomMessages, len(entRoomMessages))
	for i, entRoomMessage := range entRoomMessages {
		roomMessages[i] = FromEntRoomMessage(entRoomMessage)
	}
	return roomMessages
}

func FromEntRoomMessagesWithErrHandle(entRoomMessages []*ent.RoomMessage, err error) (model.RoomMessages, error) {
	return FromEntRoomMessages(entRoomMessages), err
}

// func FromEntUser(entUser *ent.User) *model.User {
// 	if entUser == nil {
// 		return nil
// 	}
// 	return &model.User{
// 		BaseModel: common.BaseModel{
// 			ID:        entUser.ID,
// 			CreatedAt: entUser.CreatedAt,
// 		},
// 		Username:       entUser.Username,
// 		Email:          entUser.Email,
// 		HashedPassword: entUser.HashedPassword,
// 		Fullname:       entUser.Fullname,
// 		Status:         entUser.Status,
// 	}
// }

// func FromEntUserWithErrHandle(entUser *ent.User, err error) (*model.User, error) {
// 	return FromEntUser(entUser), errhandler.HandleErr(err)
// }

// func FromEntUsers(entUsers []*ent.User) model.Users {
// 	users := make(model.Users, len(entUsers))
// 	for i, entUser := range entUsers {
// 		users[i] = FromEntUser(entUser)
// 	}
// 	return users
// }

// func FromEntUsersWithErrHandle(entUsers []*ent.User, err error) (model.Users, error) {
// 	return FromEntUsers(entUsers), errhandler.HandleErr(err)
// }
