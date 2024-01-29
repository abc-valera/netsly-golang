package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

func FromEntChatMember(entChatMember *ent.ChatMember) model.ChatMember {
	if entChatMember == nil {
		return model.ChatMember{}
	}
	return model.ChatMember{
		UserID:     entChatMember.UserID,
		ChatRoomID: entChatMember.ChatRoomID,
		CreatedAt:  entChatMember.CreatedAt,
	}
}

func FromEntChatMemberWithErrHandle(entChatMember *ent.ChatMember, err error) (model.ChatMember, error) {
	return FromEntChatMember(entChatMember), errhandler.HandleErr(err)
}

func FromEntChatMembers(entChatMembers []*ent.ChatMember) model.ChatMembers {
	chatMembers := make(model.ChatMembers, len(entChatMembers))
	for i, entChatMember := range entChatMembers {
		chatMembers[i] = FromEntChatMember(entChatMember)
	}
	return chatMembers
}

func FromEntChatMembersWithErrHandle(entChatMembers []*ent.ChatMember, err error) (model.ChatMembers, error) {
	return FromEntChatMembers(entChatMembers), errhandler.HandleErr(err)
}
