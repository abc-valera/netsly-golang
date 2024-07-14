package localFileSaverQuery

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type localFileSaver struct {
	filesPath string
}

func New(filesPath string) query.IFileContent {
	return &localFileSaver{
		filesPath: filesPath,
	}
}

func (q localFileSaver) GetByID(ctx context.Context, fileID string) ([]byte, error) {
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
