package query

import "context"

type IFileContent interface {
	GetByID(ctx context.Context, fileID string) ([]byte, error)
}
