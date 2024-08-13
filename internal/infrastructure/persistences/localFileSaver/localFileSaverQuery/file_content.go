package localFileSaverQuery

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
)

type fileContent struct {
	filesPath string
}

func New(filesPath string) query.IFileContent {
	return &fileContent{
		filesPath: filesPath,
	}
}

func (q fileContent) GetByID(ctx context.Context, fileID string) (model.FileContent, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	content, err := os.ReadFile(q.filesPath + "/" + fileID)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, model.ErrFileContentNotFound
		}
		return nil, coderr.NewInternalErr(err)
	}

	return content, nil
}
