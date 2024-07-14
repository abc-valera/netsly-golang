package localFileSaverCommand

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
)

type localFileSaver struct {
	filesPath string
}

func New(filesPath string) command.IFileContent {
	return &localFileSaver{
		filesPath: filesPath,
	}
}

func (c localFileSaver) Create(ctx context.Context, fileID string, content []byte) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	newFile, err := os.Create(c.filesPath + "/" + fileID)
	if err != nil {
		return coderr.NewInternalErr(err)
	}
	defer newFile.Close()

	if _, err := newFile.Write(content); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}

func (c localFileSaver) Update(ctx context.Context, fileID string, newContent []byte) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := c.Delete(ctx, fileID); err != nil {
		return err
	}

	if err := c.Create(ctx, fileID, newContent); err != nil {
		return err
	}

	return nil
}

func (c localFileSaver) Delete(ctx context.Context, fileID string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := os.Remove(c.filesPath + "/" + fileID); err != nil {
		if os.IsNotExist(err) {
			return model.ErrFileContentNotFound
		}
		return coderr.NewInternalErr(err)
	}

	return nil
}
