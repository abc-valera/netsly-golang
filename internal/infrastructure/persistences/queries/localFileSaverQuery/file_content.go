package localFileSaverQuery

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
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
			return model.FileContent{}, model.ErrFileContentNotFound
		}
		return model.FileContent{}, coderr.NewInternalErr(err)
	}

	return model.FileContent{
		ID:      fileID,
		Content: content,
	}, nil
}
