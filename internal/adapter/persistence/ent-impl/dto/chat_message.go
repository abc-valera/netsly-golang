package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/model/common"
)

func FromEntChatMessage(entChatMessage *ent.ChatMessage) model.ChatMessage {
	if entChatMessage == nil {
		return model.ChatMessage{}
	}
	return model.ChatMessage{
		BaseModel: common.BaseModel{
			ID:        entChatMessage.ID,
			CreatedAt: entChatMessage.CreatedAt,
		},
		UserID:     entChatMessage.UserID,
		ChatRoomID: entChatMessage.ChatRoomID,
		Text:       entChatMessage.Text,
	}
}

func FromEntChatMessageWithErrHandle(entChatMessage *ent.ChatMessage, err error) (model.ChatMessage, error) {
	return FromEntChatMessage(entChatMessage), err
}

func FromEntChatMessages(entChatMessages []*ent.ChatMessage) model.ChatMessages {
	chatMessages := make(model.ChatMessages, len(entChatMessages))
	for i, entChatMessage := range entChatMessages {
		chatMessages[i] = FromEntChatMessage(entChatMessage)
	}
	return chatMessages
}

func FromEntChatMessagesWithErrHandle(entChatMessages []*ent.ChatMessage, err error) (model.ChatMessages, error) {
	return FromEntChatMessages(entChatMessages), err
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
