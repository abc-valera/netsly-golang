package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type roomMember struct {
	commandGeneric.ICreateUpdateDelete[model.RoomMember]
}

func NewRoomMember(db bun.IDB) command.IRoomMember {
	return &roomMember{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewRoomMember),
	}
}
