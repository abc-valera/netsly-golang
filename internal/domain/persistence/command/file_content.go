package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IFileContent interface {
	Create(ctx context.Context, req model.FileContent) (model.FileContent, error)
	Update(ctx context.Context, ids model.FileContent, req FileContentUpdateRequest) (model.FileContent, error)
	Delete(ctx context.Context, req model.FileContent) error
}

type FileContentUpdateRequest struct {
	UpdatedAt time.Time

	Content []byte
}
