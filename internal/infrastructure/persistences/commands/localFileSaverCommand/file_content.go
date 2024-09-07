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

func New(filesPath string) command.IFileContent {
	return &fileContent{
		filesPath: filesPath,
	}
}

func (c fileContent) Create(ctx context.Context, req model.FileContent) (model.FileContent, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	newFile, err := os.Create(c.filesPath + "/" + req.ID)
	if err != nil {
		return model.FileContent{}, coderr.NewInternalErr(err)
	}
	defer newFile.Close()

	if _, err := newFile.Write(req.Content); err != nil {
		return model.FileContent{}, coderr.NewInternalErr(err)
	}

	return req, nil
}

func (c fileContent) Update(ctx context.Context, ids model.FileContent, req command.FileContentUpdateRequest) (model.FileContent, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := c.Delete(ctx, ids); err != nil {
		return model.FileContent{}, err
	}

	return c.Create(ctx, model.FileContent{
		ID:      ids.ID,
		Content: req.Content,
	})
}

func (c fileContent) Delete(ctx context.Context, ids model.FileContent) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := os.Remove(c.filesPath + "/" + ids.ID); err != nil {
		if os.IsNotExist(err) {
			return model.ErrFileContentNotFound
		}
		return coderr.NewInternalErr(err)
	}

	return nil
}
