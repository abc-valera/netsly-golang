package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type user struct {
	commandGeneric.ICreateUpdateDelete[model.User]
}

func NewUser(db bun.IDB) command.IUser {
	return &user{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewUser),
	}
}
