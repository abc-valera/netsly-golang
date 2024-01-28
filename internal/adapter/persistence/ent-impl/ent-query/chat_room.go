package entquery

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/chatmember"
	"github.com/abc-valera/flugo-api-golang/gen/ent/chatroom"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query/spec"
)

type chatRoomQuery struct {
	*ent.Client
}

func NewChatRoomQuery(client *ent.Client) query.IChatRoom {
	return &chatRoomQuery{
		Client: client,
	}
}

func (c chatRoomQuery) GetByID(ctx context.Context, id string) (model.ChatRoom, error) {
	return dto.FromEntChatRoomWithErrHandle(c.ChatRoom.Get(ctx, id))
}

func (c chatRoomQuery) GetByName(ctx context.Context, name string) (model.ChatRoom, error) {
	return dto.FromEntChatRoomWithErrHandle(
		c.ChatRoom.Query().
			Where(chatroom.Name(name)).
			Only(ctx),
	)
}

func (c chatRoomQuery) GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.ChatRooms, error) {
	query := c.ChatRoom.
		Query().
		Where(chatroom.HasMembersWith(chatmember.UserID(userID)))

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntChatRoomsWithErrHandle(query.All(ctx))
}

func (c chatRoomQuery) SearchAllByName(ctx context.Context, keyword string, params spec.SelectParams) (model.ChatRooms, error) {
	query := c.ChatRoom.
		Query().
		Where(chatroom.NameContains(keyword))

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntChatRoomsWithErrHandle(query.All(ctx))

}
