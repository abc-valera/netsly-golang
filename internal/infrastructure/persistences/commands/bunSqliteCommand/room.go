package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type room struct {
	commandGeneric.ICreateUpdateDelete[model.Room]
}

func NewRoom(db bun.IDB) command.IRoom {
	return &room{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewRoom),
	}
}
