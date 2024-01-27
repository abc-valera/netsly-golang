package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	entimpl "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl"
	entcommand "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-command"
	entquery "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-query"
	enttransactioneer "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/ent-transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/messaging/dummy"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/core"
	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/global"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/port/config"
	jsonrestapi "github.com/abc-valera/flugo-api-golang/internal/port/json-rest-api"
)

// services initializes all services
func services(config config.Config) core.Services {
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

	return core.NewServices(
		logger,
		emailSender,
		passwordMaker,
		tokenMaker,
		broker,
	)
}

// persistence initializes persistence
func persistence(config config.Config) (
	core.Commands,
	core.Queries,
	transactioneer.ITransactioneer,
) {
	// Init dependencies
	client := coderr.Must[*ent.Client](entimpl.InitEntClient(config.PosrgresURL))

	// Init commands
	commands := core.NewCommands(
		entcommand.NewUserCommand(client),
		entcommand.NewJokeCommand(client),
		entcommand.NewLikeCommand(client),
		entcommand.NewCommentCommand(client),
		entcommand.NewChatRoomCommand(client),
		entcommand.NewChatMemberCommand(client),
		entcommand.NewChatMessageCommand(client),
	)

	// Init queries
	queries := core.NewQueries(
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
	commands, queries, tx := persistence(config)

	// Init domains
	domains := core.NewDomains(commands, queries, services)

	// Init usecases
	usecases := application.NewUseCases(queries, tx, domains, services)

	// Init server
	server := jsonrestapi.NewServer(
		config.JsonRestApiPort,
		config.JsonRestApiStaticPath,
		queries,
		domains,
		services,
		usecases,
	)

	// Run server
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Run server error: ", err)
		}
	}()
	global.Log.Info("json-rest-api server started", "port", config.JsonRestApiPort)

	// Stop program execution until receiving an interrupt signal
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	<-gracefulShutdown

	// After receiving an interrupt signal, wait for all requests to be processed or 15 seconds
	// and then shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown server error: ", err)
	}
}
