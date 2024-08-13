package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IFileContent interface {
	GetByID(ctx context.Context, fileID string) (model.FileContent, error)
}
