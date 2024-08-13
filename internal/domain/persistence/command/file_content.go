package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type IFileContent interface {
	Create(ctx context.Context, req FileContentCreateRequest) (model.FileContent, error)
	Update(ctx context.Context, id string, req FileContentUpdateRequest) (model.FileContent, error)
	Delete(ctx context.Context, id string) error
}

type FileContentCreateRequest struct {
	ID      string
	Content model.FileContent
}

type FileContentUpdateRequest struct {
	Content model.FileContent
}
