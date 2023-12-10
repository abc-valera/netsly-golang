package transactioneer

import "context"

// ITransactioneer is an interface that defines a method to perform a transaction
type ITransactioneer interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error
}
