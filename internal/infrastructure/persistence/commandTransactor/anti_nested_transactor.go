package commandTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
)

type antiNestedTransactor struct{}

// NewAntiNested returns an instance of CommandTransactor that returns an error
// if trying to execute it.
//
// Is used for preventing nested transactions.
func NewAntiNested() commandTransactor.ITransactor {
	return &transactor{}
}

func (t antiNestedTransactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands persistence.Commands) error,
) error {
	return coderr.NewInternalString("Nested transactions are not allowed")
}
