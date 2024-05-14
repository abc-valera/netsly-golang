package transactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
)

// ITransactor is an interface that defines a method to perform a transaction
type ITransactor interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context, txEntities domain.Entities) error) error
}
