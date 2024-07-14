package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/core/telemetry"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/seed"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp"
)

func main() {
	// Init global application mode
	global.InitMode(mode.Mode(LoadEnv("MODE")))

	// Init services
	services := service.New(
		LoadEnv("LOGGER_SERVICE"),
		LoadEnv("LOGGER_SERVICE_LOGS_FOLDER_PATH"),

		LoadEnv("EMAIL_SENDER_SERVICE"),

		LoadEnv("TASK_QUEUER_SERVICE"),
	)

	// Init global logger
	global.InitLog(services.Logger)

	// Init persistence
	commands, queries, commandTransactor := persistence.New(
		LoadEnv("POSTGRES_URL"),
		LoadEnv("FILE_SAVER_FILES_PATH"),
	)

	// Init entities
	entities := domain.NewEntities(commands, commandTransactor, queries, services)

	// Init transactor
	tx := entityTransactor.New(commandTransactor, queries, services)

	// Init usecases
	usecases := application.NewUsecases(tx, entities, services)

	// Get cli flags
	entrypoint := flag.String("entrypoint", "webApp", "Port flag specifies the application presentation to be run: webApp, jsonApi, grpcApi")
	flag.Parse()

	// Init OpenTelemetry instrumentation
	jaegerTraceExporter := coderr.Must(telemetry.NewJaegerTraceExporter())

	jaegerTraceProvider := telemetry.NewTraceProvider(jaegerTraceExporter, "netsly."+*entrypoint)
	defer jaegerTraceProvider.Shutdown(context.Background())

	global.InitTracer(jaegerTraceProvider.Tracer("netsly-golang"))

	// Init server functions
	var serverStart, serverGracefulStop func()

	switch *entrypoint {
	case "webApp":
		serverStart, serverGracefulStop = webApp.NewServer(
			LoadEnv("WEB_APP_PORT"),
			LoadEnv("WEB_APP_TEMPLATE_PATH"),
			LoadEnv("STATIC_PATH"),
			services,
			entities,
			usecases,
		)
	case "jsonApi":
		serverStart, serverGracefulStop = jsonApi.NewServer(
			LoadEnv("JSON_API_PORT"),
			LoadEnv("STATIC_PATH"),
			LoadEnv("JWT_SIGN_KEY"),
			LoadEnvTime("ACCESS_TOKEN_DURATION"),
			LoadEnvTime("REFRESH_TOKEN_DURATION"),
			entities,
			services,
			usecases,
		)
	case "grpcApi":
		serverStart, serverGracefulStop = grpcApi.RunServer(
			LoadEnv("GRPC_API_PORT"),
			LoadEnv("STATIC_PATH"),
			services,
			usecases,
		)
	case "seed":
		seed.Seed(
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

func LoadEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return strings.TrimSpace(env)
}

func LoadEnvTime(key string) time.Duration {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return coderr.Must(time.ParseDuration(env))
}
