package bunSqliteCommand

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
)

type comment struct {
	commandGeneric.ICreateUpdateDelete[model.Comment]
}

func NewComment(db bun.IDB) command.IComment {
	return &comment{
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewComment),
	}
}
