package entquery

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/chatmember"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/query"
)

type chatMemberQuery struct {
	*ent.Client
}

func NewChatMemberQuery(client *ent.Client) query.IChatMember {
	return &chatMemberQuery{
		Client: client,
	}
}

func (cq chatMemberQuery) GetByIDs(ctx context.Context, userID string, chatRoomID string) (model.ChatMember, error) {
	return dto.FromEntChatMemberWithErrHandle(cq.ChatMember.Query().Where(
		chatmember.UserID(userID),
		chatmember.ChatRoomID(chatRoomID),
	).Only(ctx))
}
