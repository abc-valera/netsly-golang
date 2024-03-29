package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/core/mode"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/transactioneer"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/gqlApi"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/grpcApi"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/seed"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/webApp"
)

var (
	modeEnv = os.Getenv("MODE")

	webAppPortEnv         = os.Getenv("WEB_APP_PORT")
	webApptemplatePathEnv = os.Getenv("WEB_APP_TEMPLATE_PATH")
	webAppStaticPathEnv   = os.Getenv("WEB_APP_STATIC_PATH")

	jsonApiPortEnv       = os.Getenv("JSON_API_PORT")
	jsonApiStaticPathEnv = os.Getenv("JSON_API_STATIC_PATH")

	grpcApiPortEnv       = os.Getenv("GRPC_API_PORT")
	grpcApiStaticPathEnv = os.Getenv("GRPC_API_STATIC_PATH")
)

func main() {
	// Init global variables

	global.InitLog(slogLogger.New())

	var appMode mode.Mode
	switch modeEnv {
	case "dev":
		appMode = mode.Development
	case "prod":
		appMode = mode.Production
	default:
		coderr.Fatal("'MODE' environmental variable is invalid")
	}
	global.InitMode(appMode)

	// Init services
	services := service.Init()

	// Init persistence dependencies
	db := persistence.InitDB()

	// Init persistence
	commands, queries := persistence.InitCommands(db), persistence.InitQueries(db)

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

	// Init transaction
	tx := transactioneer.NewTransactioneer(db, services)

	// Init usecases
	usecases := application.NewUseCases(queries, tx, entities, services)

	// Get cli flags
	entrypoint := flag.String("entrypoint", "webApp", "Port flag specifies the application presentation to be run: webApp, jsonApi, grpcApi")
	flag.Parse()

	// Init server functions
	var serverStart, serverGracefulStop func()
	switch *entrypoint {
	case "webApp":
		serverStart, serverGracefulStop = webApp.NewServer(
			webAppPortEnv,
			webApptemplatePathEnv,
			webAppStaticPathEnv,
			queries,
			entities,
			services,
			usecases,
		)
	case "jsonApi":
		serverStart, serverGracefulStop = jsonApi.NewServer(
			jsonApiPortEnv,
			jsonApiStaticPathEnv,
			queries,
			entities,
			services,
			usecases,
		)
	case "grpcApi":
		serverStart, serverGracefulStop = grpcApi.RunServer(
			grpcApiPortEnv,
			grpcApiStaticPathEnv,
			services,
			usecases,
		)
	case "gqlApi":
		serverStart, serverGracefulStop = gqlApi.NewServer()
	case "seed":
		seed.Seed(
			queries,
			entities,
			tx,
		)
		return
	default:
		coderr.Fatal("Provided invalid entrypoint flag")
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
