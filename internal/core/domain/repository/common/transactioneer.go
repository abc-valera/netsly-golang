package common

import "context"

// Transactioneer is an interface that defines a method to perform a transaction
type Transactioneer interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error
}
