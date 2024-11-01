package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand/bunSqliteCommandGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel/trace"
)

type fileInfo struct {
	db bun.IDB
	commandGeneric.ICreateUpdateDelete[model.FileInfo]
}

func NewFileInfo(db bun.IDB) command.IFileInfo {
	return &fileInfo{
		db:                  db,
		ICreateUpdateDelete: bunSqliteCommandGeneric.New(db, bunSqliteDto.NewFileInfo),
	}
}

func (c fileInfo) LinkWithJoke(ctx context.Context, fileInfoID string, jokeID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	bunJunctionModel := bunSqliteDto.FileInfosJoke{
		FileInfoID: fileInfoID,
		JokeID:     jokeID,
	}
	res, err := c.db.NewInsert().Model(&bunJunctionModel).Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}

func (c fileInfo) LinkWithRoom(ctx context.Context, fileInfoID string, roomID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	bunJunctionModel := bunSqliteDto.FileInfosRoom{
		FileInfoID: fileInfoID,
		RoomID:     roomID,
	}
	res, err := c.db.NewInsert().Model(&bunJunctionModel).Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}
