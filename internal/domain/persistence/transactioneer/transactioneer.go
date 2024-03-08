package transactioneer

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
)

// ITransactioneer is an interface that defines a method to perform a transaction
type ITransactioneer interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context, commands domain.Commands) error) error
}
