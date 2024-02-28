package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type roomQuery struct {
	*ent.Client
}

func NewRoomQuery(client *ent.Client) query.IRoom {
	return &roomQuery{
		Client: client,
	}
}

func (c roomQuery) GetByID(ctx context.Context, id string) (model.Room, error) {
	return dto.FromEntRoomWithErrHandle(c.Room.Get(ctx, id))
}

func (c roomQuery) GetByName(ctx context.Context, name string) (model.Room, error) {
	return dto.FromEntRoomWithErrHandle(
		c.Room.Query().
			Where(room.Name(name)).
			Only(ctx),
	)
}

func (c roomQuery) GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Rooms, error) {
	query := c.Room.
		Query().
		Where(room.HasMembersWith(roommember.HasUserWith(user.ID(userID))))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntRoomsWithErrHandle(query.All(ctx))
}

func (c roomQuery) SearchAllByName(ctx context.Context, keyword string, params spec.SelectParams) (model.Rooms, error) {
	query := c.Room.
		Query().
		Where(room.NameContains(keyword))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntRoomsWithErrHandle(query.All(ctx))

}
