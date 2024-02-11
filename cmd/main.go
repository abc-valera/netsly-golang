package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

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
	"github.com/abc-valera/netsly-api-golang/internal/domain/seed"
	grpcapi "github.com/abc-valera/netsly-api-golang/internal/port/grpc-api"
	jsonrestapi "github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api"
	webapp "github.com/abc-valera/netsly-api-golang/internal/port/web-app"
)

var (
	modeEnv = os.Getenv("MODE")

	postgresUrlEnv = os.Getenv("POSTGRES_URL")

	accessTokenDurationEnv  = os.Getenv("ACCESS_TOKEN_DURATION")
	refreshTokenDurationEnv = os.Getenv("REFRESH_TOKEN_DURATION")
	signKeyEnv              = os.Getenv("JWT_SIGN_KEY")

	webAppPortEnv         = os.Getenv("WEB_APP_PORT")
	webApptemplatePathEnv = os.Getenv("WEB_APP_TEMPLATE_PATH")
	webAppStaticPathEnv   = os.Getenv("WEB_APP_STATIC_PATH")

	jsonRestApiPortEnv       = os.Getenv("JSON_REST_API_PORT")
	jsonRestApiStaticPathEnv = os.Getenv("JSON_REST_API_STATIC_PATH")

	grpcApiPortEnv       = os.Getenv("GRPC_API_PORT")
	grpcApiStaticPathEnv = os.Getenv("GRPC_API_STATIC_PATH")
)

func main() {
	// Init global variables
	logger := logger.NewSlogLogger()
	global.InitLog(logger)

	var appMode mode.Mode
	switch modeEnv {
	case "dev":
		appMode = mode.Development
	case "prod":
		appMode = mode.Production
	default:
		global.Log().Fatal("'MODE' environmental variable is invalid")
	}
	global.InitMode(appMode)

	// Init services
	passwordMaker := password.NewPasswordMaker()
	tokenMaker := token.NewTokenMaker(
		coderr.Must[time.Duration](time.ParseDuration(accessTokenDurationEnv)),
		coderr.Must[time.Duration](time.ParseDuration(refreshTokenDurationEnv)),
		signKeyEnv,
	)
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
	client := coderr.Must[*ent.Client](entimpl.InitEntClient(postgresUrlEnv))

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

	// Get cli flags
	entrypoint := flag.String("entrypoint", "web-app", "Port flag specifies the application port to be run: web-app, json-rest-api, grpc-api")
	flag.Parse()

	// Init server functions
	var serverStart, serverGracefulStop func()
	switch *entrypoint {
	case "web-app":
		serverStart, serverGracefulStop = webapp.NewServer(
			webAppPortEnv,
			webApptemplatePathEnv,
			webAppStaticPathEnv,
			queries,
			entities,
			services,
			usecases,
		)
	case "json-rest-api":
		serverStart, serverGracefulStop = jsonrestapi.NewServer(
			jsonRestApiPortEnv,
			jsonRestApiStaticPathEnv,
			queries,
			entities,
			services,
			usecases,
		)
	case "grpc-api":
		serverStart, serverGracefulStop = grpcapi.RunServer(
			grpcApiPortEnv,
			grpcApiStaticPathEnv,
			services,
			usecases,
		)
	case "seed":
		seed.Seed(
			queries,
			entities,
		)
		return
	default:
		global.Log().Fatal("Provided invalid entrypoint flag")
		return
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
