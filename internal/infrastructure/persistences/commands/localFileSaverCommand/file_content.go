package localFileSaverCommand

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

type fileContent struct {
	filesPath string
}

func New(filesPath string) command.ICreateUpdateDelete[model.FileContent] {
	return &fileContent{
		filesPath: filesPath,
	}
}

func (c fileContent) Create(ctx context.Context, req model.FileContent) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	newFile, err := os.Create(c.filesPath + "/" + req.ID)
	if err != nil {
		return coderr.NewInternalErr(err)
	}
	defer newFile.Close()

	if _, err := newFile.Write(req.Content); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}

func (c fileContent) Update(ctx context.Context, req model.FileContent) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := c.Delete(ctx, req); err != nil {
		return err
	}

	return c.Create(ctx, req)
}

func (c fileContent) Delete(ctx context.Context, req model.FileContent) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := os.Remove(c.filesPath + "/" + req.ID); err != nil {
		if os.IsNotExist(err) {
			return model.ErrFileContentNotFound
		}
		return coderr.NewInternalErr(err)
	}

	return nil
}
