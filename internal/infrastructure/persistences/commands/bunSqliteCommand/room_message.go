package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type roomMessage struct {
	commandGeneric.ICreateUpdateDelete[model.RoomMessage]
}

func NewRoomMessage(db bun.IDB) command.IRoomMessage {
	return &roomMessage{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewRoomMessage),
	}
}
