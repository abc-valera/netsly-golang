package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

type IFileInfo interface {
	GetByID(ctx context.Context, id string) (model.FileInfo, error)
	GetAll(ctx context.Context, selector selector.Selector) (model.FileInfos, error)
}
