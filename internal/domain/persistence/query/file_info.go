package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type IFileInfo interface {
	GetByID(ctx context.Context, id string) (model.FileInfo, error)
	GetAll(ctx context.Context) (model.FileInfos, error)
}
