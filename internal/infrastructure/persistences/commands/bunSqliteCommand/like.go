package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type like struct {
	commandGeneric.ICreateUpdateDelete[model.Like]
}

func NewLike(db bun.IDB) command.ILike {
	return &like{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewLike),
	}
}
