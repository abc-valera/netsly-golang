package localFileSaverQuery

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
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

func (q fileContent) GetOne(ctx context.Context, fitlerOptions ...filter.Option[model.FileContent]) (model.FileContent, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	filters := filter.New(fitlerOptions...)

	if len(filters) > 1 {
		return model.FileContent{}, coderr.NewInternalString("only one filter option is allowed for file content")
	}

	if len(filters) < 1 {
		return model.FileContent{}, coderr.NewInternalString("file content 'ID' filter option is required")
	}

	fileContentID := filters[0].By.ID

	content, err := os.ReadFile(q.filesPath + "/" + fileContentID)
	if err != nil {
		if os.IsNotExist(err) {
			return model.FileContent{}, model.ErrFileContentNotFound
		}
		return model.FileContent{}, coderr.NewInternalErr(err)
	}

	return model.FileContent{
		ID:      fileContentID,
		Content: content,
	}, nil
}
