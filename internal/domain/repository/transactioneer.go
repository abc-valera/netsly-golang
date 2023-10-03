package repository

import "context"

type Transactioneer interface {
	PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error
}
