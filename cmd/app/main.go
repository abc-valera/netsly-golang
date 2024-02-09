package main

import (
	"flag"
	"os"
	"os/signal"

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
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/mode"
	grpcapi "github.com/abc-valera/netsly-api-golang/internal/port/grpc-api"
	jsonrestapi "github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api"
	webapp "github.com/abc-valera/netsly-api-golang/internal/port/web-app"
)

func main() {
	// Get cli flags
	entrypoint := *flag.String("entrypoint", "web-app", "Port flag specifies the application port to be run: web-app, json-rest-api, grpc-api")

	// Init global variables
	var appMode mode.Mode
	switch os.Getenv("MODE") {
	case "dev":
		appMode = mode.Development
	case "prod":
		appMode = mode.Production
	default:
		global.Log().Fatal("'MODE' environmental variable is invalid")
	}

	logger := logger.NewSlogLogger()

	global.Init(
		appMode,
		logger,
	)

	// Init services
	passwordMaker := password.NewPasswordMaker()
	tokenMaker := token.NewTokenMaker()
	emailSender := email.NewDummyEmailSender()
	broker := dummy.NewMessagingBroker(emailSender)

	services := domain.NewServices(
		logger,
		emailSender,
		passwordMaker,
		tokenMaker,
		broker,
	)

	// Init persistence dependencies
	client := coderr.Must[*ent.Client](entimpl.InitEntClient())

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

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

	// Init usecases
	usecases := application.NewUseCases(queries, tx, entities, services)

	// Init server functions
	var serverStart, serverGracefulStop func()
	switch entrypoint {
	case "web-app":
		serverStart, serverGracefulStop = webapp.NewServer(queries, entities, services, usecases)
	case "json-rest-api":
		serverStart, serverGracefulStop = jsonrestapi.NewServer(queries, entities, services, usecases)
	case "grpc-api":
		serverStart, serverGracefulStop = grpcapi.RunServer(services, usecases)
	default:
		global.Log().Fatal("Provided invalid entrypoint flag")
	}

	// Run server
	go serverStart()

	// Stop program execution until receiving an interrupt signal
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	<-gracefulShutdown

	// After receiving an interrupt signal, run graceful stop
	serverGracefulStop()
}
