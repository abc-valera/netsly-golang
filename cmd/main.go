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
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/nopLogger"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/passworder"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/taskQueuer/dummyTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/seed"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp"
)

func main() {
	// For the configuration the environment variables are used

	// Get cli flags for the entrypoint
	entrypoint := flag.String("entrypoint", "webApp", "Port flag specifies the application presentation to be run: webApp, jsonApi, grpcApi")
	flag.Parse()

	// Init OpenTelemetry instrumentation
	jaegerTraceExporter := coderr.Must(telemetry.NewJaegerTraceExporter())

	jaegerTraceProvider := telemetry.NewTraceProvider(jaegerTraceExporter, "netsly."+*entrypoint)
	defer jaegerTraceProvider.Shutdown(context.Background())

	// Init services
	var services domain.Services

	services.Passworder = passworder.New()

	switch loggerService := LoadEnv("LOGGER_SERVICE"); loggerService {
	case "nop":
		services.Logger = nopLogger.New()
	case "slog":
		services.Logger = slogLogger.New(LoadEnv("LOGGER_SERVICE_LOGS_FOLDER_PATH"))
	default:
		coderr.Fatal("Invalid Logger implementation provided: " + loggerService)
	}

	switch emailSenderService := LoadEnv("EMAIL_SENDER_SERVICE"); emailSenderService {
	case "dummy":
		services.EmailSender = dummyEmailSender.New()
	default:
		coderr.Fatal("Invalid Email Sender implementation provided: " + emailSenderService)
	}

	switch taskQueuerService := LoadEnv("TASK_QUEUER_SERVICE"); taskQueuerService {
	case "dummy":
		services.TaskQueuer = dummyTaskQueuer.New(services.EmailSender)
	default:
		coderr.Fatal("Invalid Task Queuer implementation provided: " + taskQueuerService)
	}

	// Init global
	global.Init(
		mode.Mode(LoadEnv("MODE")),
		jaegerTraceProvider.Tracer("netsly-golang"),
		services.Logger,
	)

	// Init persistence
	var commandsAndQueriesDependencies implementation.CommandsAndQueriesDependencies
	var commandTransactorDependencies commandTransactor.Dependencies
	var entityTransactorDependencies entityTransactor.Dependencies

	if gormSqliteEnv := strings.TrimSpace(os.Getenv("GORM_SQLITE_PATH")); gormSqliteEnv != "" {
		gormSqliteDependency := coderr.Must(gormSqlite.New(gormSqliteEnv))

		commandsAndQueriesDependencies.GormSqlite = gormSqliteDependency
		commandTransactorDependencies.GormSqlite = gormSqliteDependency
		entityTransactorDependencies.GormSqlite = gormSqliteDependency
	}

	if boilerSqliteEnv := strings.TrimSpace(os.Getenv("BOILER_SQLITE_PATH")); boilerSqliteEnv != "" {
		boilerSqliteDependency := coderr.Must(boilerSqlite.New(boilerSqliteEnv))

		commandsAndQueriesDependencies.BoilerSqlite = boilerSqliteDependency
		commandTransactorDependencies.BoilerSqlite = boilerSqliteDependency
		entityTransactorDependencies.BoilerSqlite = boilerSqliteDependency
	}

	if localFileSaverEnv := strings.TrimSpace(os.Getenv("LOCAL_FILE_SAVER_FILES_PATH")); localFileSaverEnv != "" {
		localFileSaverDependency := coderr.Must(localFileSaver.New(localFileSaverEnv))

		commandsAndQueriesDependencies.LocalFileSaver = localFileSaverDependency
		commandTransactorDependencies.LocalFileSaver = localFileSaverDependency
		entityTransactorDependencies.LocalFileSaver = localFileSaverDependency
	}

	commands, queries, err := implementation.NewCommandsAndQueries(commandsAndQueriesDependencies)
	if err != nil {
		coderr.Fatal(err)
	}

	// Init command transactor
	commandTransactor := commandTransactor.New(commandTransactorDependencies)

	// Init entities
	entities := domain.NewEntities(commands, commandTransactor, queries, services)

	// Init Entity Transactor
	entityTransactor := entityTransactor.New(entityTransactorDependencies, services)

	// Init usecases
	usecases := application.NewUsecases(entityTransactor, entities, services)

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

// LoadEnv is a shortcut for trimming and empty-cheking environemnt variables.
// Stops the program execution if variable is not set.
func LoadEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return strings.TrimSpace(env)
}

// LoadEnv is a shortcut for loading and parsing environment variables as time.Duration.
// Stops the program execution if variable is not set or parsing error occurs.
func LoadEnvTime(key string) time.Duration {
	env := os.Getenv(key)
	if env == "" {
		coderr.Fatal(key + " environment variable is not set")
	}

	return coderr.Must(time.ParseDuration(env))
}
