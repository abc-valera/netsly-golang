package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IFileInfo interface {
	Create(ctx context.Context, req model.FileInfo) (model.FileInfo, error)
	Update(ctx context.Context, id string, req FileInfoUpdateRequest) (model.FileInfo, error)
	Delete(ctx context.Context, id string) error
}

type FileInfoUpdateRequest struct {
	Name *string
}
