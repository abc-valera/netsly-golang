package query

import (
	"context"
)

type IFileContent interface {
	Get(ctx context.Context, id string) ([]byte, error)
}
