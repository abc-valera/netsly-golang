package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IFileInfo interface {
	Create(ctx context.Context, req model.FileInfo) (model.FileInfo, error)
	Update(ctx context.Context, ids model.FileInfo, req FileInfoUpdateRequest) (model.FileInfo, error)
	Delete(ctx context.Context, req model.FileInfo) error
}

type FileInfoUpdateRequest struct {
	UpdatedAt time.Time

	Name *string
}
