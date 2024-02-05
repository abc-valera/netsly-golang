package main

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	entimpl "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl"
	entcommand "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/ent-command"
	entquery "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/ent-query"
	enttransactioneer "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/ent-transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/messaging/dummy"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/domain/seed"
	"github.com/abc-valera/netsly-api-golang/internal/port/config"
)

func services(config config.Config) domain.Services {
	emailSender :=
		email.NewDummyEmailSender()
	logger :=
		logger.NewSlogLogger()
	passwordMaker :=
		password.NewPasswordMaker()
	tokenMaker :=
		token.NewTokenMaker(config.AccessTokenDuration, config.RefreshTokenDuration)
	broker :=
		dummy.NewMessagingBroker(emailSender)

	return domain.NewServices(
		logger,
		emailSender,
		passwordMaker,
		tokenMaker,
		broker,
	)
}

func persistence(config config.Config) (
	domain.Commands,
	domain.Queries,
	transactioneer.ITransactioneer,
) {
	// Init dependencies
	client := coderr.Must[*ent.Client](entimpl.InitEntClient(config.PosrgresURL))

	// Init commands
	commands := domain.NewCommands(
		entcommand.NewUserCommand(client),
		entcommand.NewJokeCommand(client),
		entcommand.NewLikeCommand(client),
		entcommand.NewCommentCommand(client),
		entcommand.NewChatRoomCommand(client),
		entcommand.NewChatMemberCommand(client),
		entcommand.NewChatMessageCommand(client),
	)

	// Init queries
	queries := domain.NewQueries(
		entquery.NewUserQuery(client),
		entquery.NewJokeQuery(client),
		entquery.NewLikeQuery(client),
		entquery.NewCommentQuery(client),
		entquery.NewChatRoomQuery(client),
		entquery.NewChatMemberQuery(client),
		entquery.NewChatMessageQuery(client),
	)

	// Init transactioneer
	tx := enttransactioneer.NewTransactioneer(client)

	return commands, queries, tx
}

func main() {
	// Init config
	config := coderr.Must[config.Config](config.NewConfig())

	// Init services
	services := services(config)

	// Init persistence
	commands, queries, _ := persistence(config)

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

	// Init seed data
	seed.Seed(queries, entities)
}
