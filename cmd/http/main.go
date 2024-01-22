package main

import (
	"log"
	"os"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/config"
	entimpl "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl"
	entcommand "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-command"
	entquery "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-query"
	enttransactioneer "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/messaging/dummy"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http"
)

// services initializes all services
func services(config config.Config) service.Services {
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

	return service.NewServices(
		logger,
		emailSender,
		passwordMaker,
		tokenMaker,
		broker,
	)
}

// persistence initializes persistence
func persistence(config config.Config) (
	command.Commands,
	query.Queries,
	transactioneer.ITransactioneer,
) {
	// Init dependencies
	client := coderr.Must[*ent.Client](entimpl.InitEntClient(config.DatabaseURL))

	// Init commands
	commands := command.NewCommands(
		entcommand.NewUserCommand(client),
		entcommand.NewJokeCommand(client),
		entcommand.NewLikeCommand(client),
		entcommand.NewCommentCommand(client),
		entcommand.NewChatRoomCommand(client),
		entcommand.NewChatMemberCommand(client),
		entcommand.NewChatMessageCommand(client),
	)

	// Init queries
	queries := query.NewQueries(
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
	config := coderr.Must[config.Config](config.NewConfig(os.Getenv("CONFIG_PATH")))

	// Init services
	services := services(config)

	// Init persistence
	commands, queries, tx := persistence(config)

	// Init domains
	domains := domain.NewDomains(commands, queries, services)

	// Init usecases
	usecases := application.NewUseCases(queries, tx, domains, services)

	if err := http.RunServer(
		config.HTMXPort,
		config.TemplatePath,
		queries,
		domains,
		services,
		usecases,
	); err != nil {
		log.Fatal("Run server error: ", err)
	}
}
