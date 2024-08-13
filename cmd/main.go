package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/env"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/core/opentelemetry"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/infrastructure"
	"github.com/abc-valera/netsly-golang/internal/presentation/grpcApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/seed"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp"
)

func main() {
	// For the configuration the environment variables are used

	// Get cli flags for the entrypoint
	entrypoint := flag.String("entrypoint", "webApp", "Port flag specifies the application presentation to be run: webApp, jsonApi, grpcApi")
	flag.Parse()

	// Init Mode first
	global.InitMode()

	// Init services
	services := infrastructure.NewServices()

	// Init other global variables
	global.InitLog(services.Logger)
	global.InitTracer(coderr.Must(
		opentelemetry.NewTracer(services.OtelTraceExporter, "netsly."+*entrypoint),
	))

	// Init persistence
	commands, queries, commandTransactor, entityTransactor := infrastructure.NewPersistences(services)

	// Init entities
	entities := domain.NewEntities(commands, commandTransactor, queries, services)

	// Init usecases
	usecases := application.NewUsecases(entityTransactor, entities, services)

	// Init server functions
	var serverStart, serverGracefulStop func()

	switch *entrypoint {
	case "webApp":
		serverStart, serverGracefulStop = webApp.NewServer(
			env.Load("WEB_APP_PORT"),
			env.Load("WEB_APP_TEMPLATE_PATH"),
			env.Load("STATIC_PATH"),
			services,
			entities,
			usecases,
		)
	case "jsonApi":
		serverStart, serverGracefulStop = jsonApi.NewServer(
			env.Load("JSON_API_PORT"),
			env.Load("JWT_SIGN_KEY"),
			env.LoadDuration("ACCESS_TOKEN_DURATION"),
			env.LoadDuration("REFRESH_TOKEN_DURATION"),
			entities,
			services,
			usecases,
		)
	case "grpcApi":
		serverStart, serverGracefulStop = grpcApi.RunServer(
			env.Load("GRPC_API_PORT"),
			env.Load("STATIC_PATH"),
			services,
			usecases,
		)
	case "seed":
		seed.Seed(
			entityTransactor,
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
