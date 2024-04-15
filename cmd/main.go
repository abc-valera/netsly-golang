package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/core/mode"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/logger/nopLogger"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/passwordMaker"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/taskQueuer/dummyTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/service/tokenMaker"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/transactor"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/grpcApi"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/seed"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/webApp"
)

// Load environment variables
var (
	modeEnv = LoadEnv("MODE")

	loggerServiceEnv = LoadEnv("LOGGER_SERVICE")
	emailSenderEnv   = LoadEnv("EMAIL_SENDER_SERVICE")
	taskQueuerEnv    = LoadEnv("TASK_QUEUER_SERVICE")

	accessTokenDurationEnv  = LoadEnvTime("ACCESS_TOKEN_DURATION")
	refreshTokenDurationEnv = LoadEnvTime("REFRESH_TOKEN_DURATION")
	signKeyEnv              = LoadEnv("JWT_SIGN_KEY")

	postgresUrlEnv = LoadEnv("POSTGRES_URL")

	webAppPortEnv         = os.Getenv("WEB_APP_PORT")
	webApptemplatePathEnv = os.Getenv("WEB_APP_TEMPLATE_PATH")
	webAppStaticPathEnv   = os.Getenv("WEB_APP_STATIC_PATH")

	jsonApiPortEnv       = os.Getenv("JSON_API_PORT")
	jsonApiStaticPathEnv = os.Getenv("JSON_API_STATIC_PATH")

	grpcApiPortEnv       = os.Getenv("GRPC_API_PORT")
	grpcApiStaticPathEnv = os.Getenv("GRPC_API_STATIC_PATH")
)

func main() {
	var Mode mode.Mode
	switch modeEnv {
	case "dev":
		Mode = mode.Development
	case "prod":
		Mode = mode.Production
	default:
		coderr.Fatal("'MODE' environmental variable is invalid")
	}

	var Log service.ILogger
	switch loggerServiceEnv {
	case "slog":
		Log = slogLogger.New()
	case "nop":
		Log = nopLogger.New()
	default:
		coderr.Fatal("'LOGGER_SERVICE' environmental variable is invalid")
	}

	// Init global variables
	global.Init(
		Mode,
		Log,
	)

	// Init services
	var EmailSender service.IEmailSender
	switch emailSenderEnv {
	case "dummy":
		EmailSender = dummyEmailSender.New()
	default:
		coderr.Fatal("EMAIL_SENDER_SERVICE environmental variable is invalid")
	}

	var TaskQueuer service.ITaskQueuer
	switch taskQueuerEnv {
	case "dummy":
		TaskQueuer = dummyTaskQueuer.New(EmailSender)
	default:
		coderr.Fatal("TASK_QUEUER_SERVICE environmental variable is invalid")
	}

	services := domain.NewServices(
		EmailSender,
		passwordMaker.New(),
		tokenMaker.NewJWT(accessTokenDurationEnv, refreshTokenDurationEnv, signKeyEnv),
		TaskQueuer,
	)

	// Init persistence dependencies
	db := persistence.InitDB(postgresUrlEnv)

	// Init persistence
	commands, queries := persistence.InitCommands(db), persistence.InitQueries(db)

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

	// Init transaction
	tx := transactor.NewTransactor(db, services)

	// Init usecases
	usecases := application.NewUseCases(tx, entities, services)

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
			services,
			entities,
			usecases,
		)
	case "jsonApi":
		serverStart, serverGracefulStop = jsonApi.NewServer(
			jsonApiPortEnv,
			jsonApiStaticPathEnv,
			services,
			entities,
			usecases,
		)
	case "grpcApi":
		serverStart, serverGracefulStop = grpcApi.RunServer(
			grpcApiPortEnv,
			grpcApiStaticPathEnv,
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
