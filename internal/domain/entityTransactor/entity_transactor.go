package entityTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
)

// ITransactor is an interface that defines a method to perform a transaction with entities.
// Should be used only inside the Application layer!
type ITransactor interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context, txEntities domain.Entities) error) error
}
