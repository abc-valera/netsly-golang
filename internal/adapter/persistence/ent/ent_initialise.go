package ent

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	entCommand "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/command"
	entQuery "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/query"
	entTx "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	domainCommand "github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	domainQuery "github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	domainTx "github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/transactioneer"

	_ "github.com/lib/pq"
)

func NewEntCommandsQueries(databaseURL string) (
	domainCommand.Commands,
	domainQuery.Queries,
	domainTx.ITransactioneer,
	error,
) {
	var commands domainCommand.Commands
	var queries domainQuery.Queries
	var tx domainTx.ITransactioneer

	// Connect to the database
	client, err := ent.Open(
		"postgres",
		databaseURL,
	)
	if err != nil {
		return commands, queries, tx, codeerr.NewInternal("NewEntImplementation", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return commands, queries, tx, codeerr.NewInternal("NewEntImplementation", err)
	}

	commands.User = entCommand.NewUserCommand(client)
	commands.Joke = entCommand.NewJokeCommand(client)
	commands.Comment = entCommand.NewCommentCommand(client)
	commands.Like = entCommand.NewLikeCommand(client)

	queries.User = entQuery.NewUserQuery(client)
	queries.Joke = entQuery.NewJokeQuery(client)
	queries.Comment = entQuery.NewCommentQuery(client)
	queries.Like = entQuery.NewLikeQuery(client)

	tx = entTx.NewTransactioneer(client)

	return commands, queries, tx, nil
}
