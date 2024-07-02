package persistence

import (
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerQuery"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PeristenceDependencies struct {
	Boiler *sql.DB
}

func NewDependencies(postgresUrl string) PeristenceDependencies {
	return PeristenceDependencies{
		Boiler: coderr.Must(boiler.Init(postgresUrl)),
	}
}

type CommandsDependencies struct {
	Boiler boil.ContextExecutor
}

func NewCommands(deps CommandsDependencies) persistence.Commands {
	return persistence.Commands{
		User:        boilerCommand.NewUser(deps.Boiler),
		Joke:        boilerCommand.NewJoke(deps.Boiler),
		Like:        boilerCommand.NewLike(deps.Boiler),
		Comment:     boilerCommand.NewComment(deps.Boiler),
		Room:        boilerCommand.NewRoom(deps.Boiler),
		RoomMember:  boilerCommand.NewRoomMember(deps.Boiler),
		RoomMessage: boilerCommand.NewRoomMessage(deps.Boiler),
		FileInfo:    boilerCommand.NewFileInfo(deps.Boiler),
	}
}

type QueriesDependencies struct {
	Boiler boil.ContextExecutor
}

func NewQueries(deps QueriesDependencies) persistence.Queries {
	return persistence.Queries{
		User:        boilerQuery.NewUser(deps.Boiler),
		Joke:        boilerQuery.NewJoke(deps.Boiler),
		Like:        boilerQuery.NewLike(deps.Boiler),
		Comment:     boilerQuery.NewComment(deps.Boiler),
		Room:        boilerQuery.NewRoom(deps.Boiler),
		RoomMember:  boilerQuery.NewRoomMember(deps.Boiler),
		RoomMessage: boilerQuery.NewRoomMessage(deps.Boiler),
		FileInfo:    boilerQuery.NewFileInfo(deps.Boiler),
	}
}
