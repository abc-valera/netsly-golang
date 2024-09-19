package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/env"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services"
	"github.com/abc-valera/netsly-golang/internal/presentation/grpcApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/seed"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp"
)

func main() {
	// Get cli flags for the entrypoint
	entrypoint := flag.String("entrypoint", "webApp", "Port flag specifies the application presentation to be run: webApp, jsonApi, grpcApi")
	flag.Parse()

	// Init Globals
	globals.New(*entrypoint)

	// Init Services
	services := services.NewServices()

	// Init DB
	db := persistences.NewDB()

	// Init Entities
	entities := entity.NewEntities(entity.NewDependency(db, services))

	// Init usecases
	usecases := application.NewUsecases(application.NewDependency(db, services))

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
		seed.Seed(entities)
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
