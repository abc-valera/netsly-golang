package commandTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
)

// ITransactor is an interface that defines a method to perform a transaction with commands.
//
// # Should be used only inside the Entity layer!
type ITransactor interface {
	PerformTX(
		ctx context.Context,
		txFunc func(ctx context.Context, txCommands persistence.Commands, txQueries persistence.Queries) error,
	) error
}
