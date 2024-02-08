package main

import (
	"os"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	entimpl "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl"
	entcommand "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/ent-command"
	entquery "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/ent-query"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/email"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/logger"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/messaging/dummy"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/password"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/service/token"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain/seed"
)

func main() {
	// Init services
	logger := logger.NewSlogLogger()
	passwordMaker := password.NewPasswordMaker()
	tokenMaker := token.NewTokenMaker()
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
	client := coderr.Must[*ent.Client](entimpl.InitEntClient())

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

	// Init global variables
	var appMode mode.Mode
	switch os.Getenv("MODE") {
	case "dev":
		appMode = mode.Development
	case "prod":
		appMode = mode.Production
	default:
		coderr.Fatal("'MODE' environmental variable is invalid")
	}

	global.Init(
		appMode,
		logger,
	)

	// Init entities
	entities := domain.NewEntities(commands, queries, services)

	// Init seed
	seed.Seed(queries, entities)
}
