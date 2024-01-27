package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
)

var (
	ErrChatRoomIDInvalid          = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid chat room ID")
	ErrChatRoomNameInvalid        = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid chat room name")
	ErrChatRoomDescriptionInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid chat room description")
)

type ChatRoom struct {
	command command.IChatRoom
}

func NewChatRoom(
	command command.IChatRoom,
) ChatRoom {
	return ChatRoom{
		command: command,
	}
}

type ChatRoomCreateRequest struct {
	Name        string
	Description string
}

func (c ChatRoom) Create(ctx context.Context, req ChatRoomCreateRequest) error {
	// Validation
	if req.Name == "" || len(req.Name) < 4 || len(req.Name) > 64 {
		return ErrChatRoomNameInvalid
	}
	if len(req.Description) > 64 {
		return ErrChatRoomDescriptionInvalid
	}

	baseModel := common.NewBaseModel()

	return c.command.Create(ctx, model.ChatRoom{
		BaseModel:   baseModel,
		Name:        req.Name,
		Description: req.Description,
	})
}

type ChatRoomUpdateRequest struct {
	Name        *string
	Description *string
}

func (c ChatRoom) Update(ctx context.Context, chatRoomID string, req ChatRoomUpdateRequest) error {
	// Validation
	if chatRoomID == "" {
		return ErrChatRoomIDInvalid
	}
	if req.Name != nil && (len(*req.Name) < 4 || len(*req.Name) > 64) {
		return ErrChatRoomNameInvalid
	}
	if req.Description != nil && len(*req.Description) > 64 {
		return ErrChatRoomDescriptionInvalid
	}

	return c.command.Update(ctx, chatRoomID, command.ChatRoomUpdate{
		Description: req.Description,
	})
}

func (c ChatRoom) Delete(ctx context.Context, chatRoomID string) error {
	// Validation
	if chatRoomID == "" {
		return ErrChatRoomIDInvalid
	}

	return c.command.Delete(ctx, chatRoomID)
}
