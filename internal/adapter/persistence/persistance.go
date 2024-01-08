package persistence

import (
	entCommand "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/command"
	entQuery "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/query"
	entTx "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/transactioneer"

	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent"
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
	client, err := ent.InitEntClient(databaseURL)
	if err != nil {
		return command.Commands{}, query.Queries{}, nil, err
	}

	commands, err := command.NewCommands(
		entCommand.NewUserCommand(client),
		entCommand.NewJokeCommand(client),
		entCommand.NewLikeCommand(client),
		entCommand.NewCommentCommand(client),
		entCommand.NewChatRoomCommand(client),
		entCommand.NewChatMemberCommand(client),
		entCommand.NewChatMessageCommand(client),
	)
	if err != nil {
		return command.Commands{}, query.Queries{}, nil, err
	}

	queries, err := query.NewQueries(
		entQuery.NewUserQuery(client),
		entQuery.NewJokeQuery(client),
		entQuery.NewLikeQuery(client),
		entQuery.NewCommentQuery(client),
		entQuery.NewChatRoomQuery(client),
		entQuery.NewChatMemberQuery(client),
		entQuery.NewChatMessageQuery(client),
	)
	if err != nil {
		return command.Commands{}, query.Queries{}, nil, err
	}

	tx := entTx.NewTransactioneer(client)

	return commands, queries, tx, nil
}
