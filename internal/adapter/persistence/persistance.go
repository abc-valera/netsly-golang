package persistence

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	entCommand "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/command"
	entQuery "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/query"
	entTx "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/transactioneer"

	entImpl "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/transactioneer"
)

func NewCommandsQueriesTx(databaseURL string) (
	command.Commands,
	query.Queries,
	transactioneer.ITransactioneer,
	error,
) {
	// Init dependencies
	client := coderr.Must[*ent.Client](entImpl.InitEntClient(databaseURL))

	// Init commands
	commands := coderr.Must[command.Commands](command.NewCommands(
		entCommand.NewUserCommand(client),
		entCommand.NewJokeCommand(client),
		entCommand.NewLikeCommand(client),
		entCommand.NewCommentCommand(client),
		entCommand.NewChatRoomCommand(client),
		entCommand.NewChatMemberCommand(client),
		entCommand.NewChatMessageCommand(client),
	))

	// Init queries
	queries := coderr.Must[query.Queries](query.NewQueries(
		entQuery.NewUserQuery(client),
		entQuery.NewJokeQuery(client),
		entQuery.NewLikeQuery(client),
		entQuery.NewCommentQuery(client),
		entQuery.NewChatRoomQuery(client),
		entQuery.NewChatMemberQuery(client),
		entQuery.NewChatMessageQuery(client),
	))

	// Init transactioneer
	tx := entTx.NewTransactioneer(client)

	return commands, queries, tx, nil
}
