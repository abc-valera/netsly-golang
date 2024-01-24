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
	webapp "github.com/abc-valera/flugo-api-golang/internal/port/web-app"
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
	client := coderr.Must[*ent.Client](entimpl.InitEntClient(config.PosrgresURL))

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

	// Init server
	server := webapp.NewServer(
		config.WebAppPort,
		config.WebAppTemplatePath,
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
	service.Log.Info("web-app server started", "port", config.WebAppPort)

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
