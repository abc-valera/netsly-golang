package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain/seed"
	sqlboilerimpl "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl"
	sqlboilercommand "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerCommand"
	sqlboilerquery "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerQuery"
	sqlboilertransactioneer "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerTransactioneer"
	grpcApi "github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi"
	jsonApi "github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi"
	webApp "github.com/abc-valera/netsly-api-golang/internal/presentation/webApp"
	"github.com/abc-valera/netsly-api-golang/internal/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/internal/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/internal/service/passwordMaker/argonPasswordMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/taskQueuer/dummyTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/service/timeMaker/baseTimeMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/tokenMaker/jwtTokenMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/uuidMaker/googleUuidMaker"
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

	jsonApiPortEnv       = os.Getenv("JSON_API_PORT")
	jsonApiStaticPathEnv = os.Getenv("JSON_API_STATIC_PATH")

	grpcApiPortEnv       = os.Getenv("GRPC_API_PORT")
	grpcApiStaticPathEnv = os.Getenv("GRPC_API_STATIC_PATH")
)

func main() {
	// Init global variables
	logger := slogLogger.New()
	global.InitLog(logger)

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
	timeMaker := baseTimeMaker.New()
	uuidMaker := googleUuidMaker.New()
	passwordMaker := argonPasswordMaker.New()
	tokenMaker := jwtTokenMaker.NewJWT(
		coderr.MustWithVal(time.ParseDuration(accessTokenDurationEnv)),
		coderr.MustWithVal(time.ParseDuration(refreshTokenDurationEnv)),
		signKeyEnv,
	)
	emailSender := dummyEmailSender.New()
	broker := dummyTaskQueuer.New(emailSender)

	services := domain.NewServices(
		logger,
		uuidMaker,
		timeMaker,
		emailSender,
		passwordMaker,
		tokenMaker,
		broker,
	)

	// Init persistence dependencies
	conn := coderr.MustWithVal(sqlboilerimpl.Init(postgresUrlEnv))

	commands := sqlboilercommand.NewCommands(conn)
	queries := sqlboilerquery.NewQueries(conn)
	tx := sqlboilertransactioneer.NewTransactioneer(conn)

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

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
