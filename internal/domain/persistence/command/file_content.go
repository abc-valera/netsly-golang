package command

import (
	"context"
)

type IFileContent interface {
	Create(ctx context.Context, fileID string, content []byte) error
	Update(ctx context.Context, fileID string, newContent []byte) error
	Delete(ctx context.Context, fileID string) error
}
