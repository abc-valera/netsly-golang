package persistence

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
)

// IDB is an abstraction for database operations.
type IDB interface {
	BeginTX(ctx context.Context) (ITX, error)
	RunInTX(
		ctx context.Context,
		fn func(context.Context, command.Commands, query.Queries) error,
	) error

	Commands() command.Commands
	Queries() query.Queries
}

// ITX is an abstraction for database transaction.
type ITX interface {
	IDB

	Commit() error
	Rollback() error
}
