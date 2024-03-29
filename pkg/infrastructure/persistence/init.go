package persistence

import (
	"database/sql"
	"os"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerCommand"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InitDB() *sql.DB {
	// Get the environment variables
	postgresUrlEnv := os.Getenv("POSTGRES_URL")
	if postgresUrlEnv != "" {
		coderr.Fatal("POSTGRES_URL environmental variable is not set")
	}

	return coderr.Must(boiler.Init(postgresUrlEnv))
}

func InitCommands(sqlboilerExecutor boil.ContextExecutor) domain.Commands {
	return domain.NewCommands(
		boilerCommand.NewUser(sqlboilerExecutor),
		boilerCommand.NewJoke(sqlboilerExecutor),
		boilerCommand.NewLike(sqlboilerExecutor),
		boilerCommand.NewComment(sqlboilerExecutor),
		boilerCommand.NewRoom(sqlboilerExecutor),
		boilerCommand.NewRoomMember(sqlboilerExecutor),
		boilerCommand.NewRoomMessage(sqlboilerExecutor),
	)
}

func InitQueries(sqlboilerExecutor boil.ContextExecutor) domain.Queries {
	return domain.NewQueries(
		boilerQuery.NewUser(sqlboilerExecutor),
		boilerQuery.NewJoke(sqlboilerExecutor),
		boilerQuery.NewLike(sqlboilerExecutor),
		boilerQuery.NewComment(sqlboilerExecutor),
		boilerQuery.NewRoom(sqlboilerExecutor),
		boilerQuery.NewRoomMember(sqlboilerExecutor),
		boilerQuery.NewRoomMessage(sqlboilerExecutor),
	)
}
