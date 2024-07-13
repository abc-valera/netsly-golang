package implementation

import (
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerQuery"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PersistenceDependencies struct {
	BoilerDB *sql.DB
}

func NewPersistenceDependencies(postgresUrl string) PersistenceDependencies {
	return PersistenceDependencies{
		BoilerDB: coderr.Must(boiler.Init(postgresUrl)),
	}
}

func NewCommands(boilerIDB boil.ContextExecutor) persistence.Commands {
	return persistence.Commands{
		User:        boilerCommand.NewUser(boilerIDB),
		Joke:        boilerCommand.NewJoke(boilerIDB),
		Like:        boilerCommand.NewLike(boilerIDB),
		Comment:     boilerCommand.NewComment(boilerIDB),
		Room:        boilerCommand.NewRoom(boilerIDB),
		RoomMember:  boilerCommand.NewRoomMember(boilerIDB),
		RoomMessage: boilerCommand.NewRoomMessage(boilerIDB),
		FileInfo:    boilerCommand.NewFileInfo(boilerIDB),
	}
}

func NewQueries(boilerIDB boil.ContextExecutor) persistence.Queries {
	return persistence.Queries{
		User:        boilerQuery.NewUser(boilerIDB),
		Joke:        boilerQuery.NewJoke(boilerIDB),
		Like:        boilerQuery.NewLike(boilerIDB),
		Comment:     boilerQuery.NewComment(boilerIDB),
		Room:        boilerQuery.NewRoom(boilerIDB),
		RoomMember:  boilerQuery.NewRoomMember(boilerIDB),
		RoomMessage: boilerQuery.NewRoomMessage(boilerIDB),
		FileInfo:    boilerQuery.NewFileInfo(boilerIDB),
	}
}
