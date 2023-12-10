package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrChatRoomIDInvalid          = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid chat room ID")
	ErrChatRoomNameInvalid        = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid chat room name")
	ErrChatRoomDescriptionInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid chat room description")
)

type ChatRoomDomain struct {
	command command.IChatRoomCommand
}

func NewChatRoomDomain(
	command command.IChatRoomCommand,
) ChatRoomDomain {
	return ChatRoomDomain{
		command: command,
	}
}

type ChatRoomCreateRequest struct {
	Name        string
	Description string
}

func (c ChatRoomDomain) Create(ctx context.Context, req ChatRoomCreateRequest) error {
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

func (c ChatRoomDomain) Update(ctx context.Context, chatRoomID string, req ChatRoomUpdateRequest) error {
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

	return c.command.Update(ctx, chatRoomID, model.ChatRoomUpdate{
		Description: req.Description,
	})
}

func (c ChatRoomDomain) Delete(ctx context.Context, chatRoomID string) error {
	// Validation
	if chatRoomID == "" {
		return ErrChatRoomIDInvalid
	}

	return c.command.Delete(ctx, chatRoomID)
}
