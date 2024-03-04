package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/errors"
)

type roomMemberCommand struct {
	*ent.Client
}

func NewRoomMemberCommand(client *ent.Client) command.IRoomMember {
	return &roomMemberCommand{
		Client: client,
	}
}

func (cm roomMemberCommand) Create(ctx context.Context, req model.RoomMember) (model.RoomMember, error) {
	roomMember, err := cm.RoomMember.Create().
		SetRoomID(req.RoomID).
		SetUserID(req.UserID).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntRoomMember(roomMember), errhandler.HandleErr(err)
}

func (cm roomMemberCommand) Delete(ctx context.Context, RoomID string, UserID string) error {
	_, err := cm.RoomMember.Delete().
		Where(
			roommember.HasRoomWith(room.ID(RoomID)),
			roommember.HasUserWith(user.ID(UserID)),
		).Exec(ctx)
	return errhandler.HandleErr(err)
}
