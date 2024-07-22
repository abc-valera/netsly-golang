package localFileSaverCommand

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
)

type fileContent struct {
	filesPath string
}

func New(filesPath string) command.IFileContent {
	return &fileContent{
		filesPath: filesPath,
	}
}

func (c fileContent) Create(ctx context.Context, req command.FileContentCreateRequest) (model.FileContent, error) {
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

	return req.Content, nil
}

func (c fileContent) Update(ctx context.Context, id string, req command.FileContentUpdateRequest) (model.FileContent, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := c.Delete(ctx, id); err != nil {
		return model.FileContent{}, err
	}

	return c.Create(ctx, command.FileContentCreateRequest{
		ID:      id,
		Content: req.Content,
	})
}

func (c fileContent) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	if err := os.Remove(c.filesPath + "/" + id); err != nil {
		if os.IsNotExist(err) {
			return model.ErrFileContentNotFound
		}
		return coderr.NewInternalErr(err)
	}

	return nil
}
