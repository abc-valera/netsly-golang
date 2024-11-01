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

func (q fileContent) Get(ctx context.Context, id string) ([]byte, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	content, err := os.ReadFile(q.filesPath + "/" + id)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, model.ErrFileContentNotFound
		}
		return nil, coderr.NewInternalErr(err)
	}

	return content, nil
}
