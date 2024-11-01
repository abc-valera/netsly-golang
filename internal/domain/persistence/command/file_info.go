package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
)

type IFileInfo interface {
	LinkWithJoke(ctx context.Context, fileInfoID, jokeID string) error
	LinkWithRoom(ctx context.Context, fileInfoID, roomID string) error
	commandGeneric.ICreateUpdateDelete[model.FileInfo]
}
