package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/chatmessage"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type chatMessageQuery struct {
	*ent.Client
}

func NewChatMessageQuery(client *ent.Client) query.IChatMessage {
	return &chatMessageQuery{
		Client: client,
	}
}

func (cm chatMessageQuery) GetByID(ctx context.Context, id string) (model.ChatMessage, error) {
	return dto.FromEntChatMessageWithErrHandle(cm.ChatMessage.Get(ctx, id))
}

func (cm chatMessageQuery) GetAllByChatRoomID(ctx context.Context, chatRoomID string) (model.ChatMessages, error) {
	return dto.FromEntChatMessagesWithErrHandle(cm.ChatMessage.Query().Where(
		chatmessage.ChatRoomID(chatRoomID),
	).All(ctx))
}

func (cm chatMessageQuery) SearchAllByText(ctx context.Context, keyword string, spec spec.SelectParams) (model.ChatMessages, error) {
	query := cm.ChatMessage.
		Query().
		Where(chatmessage.TextContains(keyword))

	if spec.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(spec.Limit)
	query.Offset(spec.Offset)

	return dto.FromEntChatMessagesWithErrHandle(query.All(ctx))
}
