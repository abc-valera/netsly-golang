package implementation

import (
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerQuery"
	localFileSaverQuery "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver/localFileSaverCommand"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PersistenceDependencies struct {
	BoilerDB  *sql.DB
	FilesPath string
}

func NewPersistenceDependencies(
	postgresUrl string,
	filesPath string,
) PersistenceDependencies {
	return PersistenceDependencies{
		BoilerDB:  coderr.Must(boiler.New(postgresUrl)),
		FilesPath: filesPath,
	}
}

func NewCommands(
	boilerIDB boil.ContextExecutor,
	filesPath string,
) persistence.Commands {
	return persistence.Commands{
		User:        boilerCommand.NewUser(boilerIDB),
		Joke:        boilerCommand.NewJoke(boilerIDB),
		Like:        boilerCommand.NewLike(boilerIDB),
		Comment:     boilerCommand.NewComment(boilerIDB),
		Room:        boilerCommand.NewRoom(boilerIDB),
		RoomMember:  boilerCommand.NewRoomMember(boilerIDB),
		RoomMessage: boilerCommand.NewRoomMessage(boilerIDB),
		FileInfo:    boilerCommand.NewFileInfo(boilerIDB),
		FileContent: localFileSaverCommand.New(filesPath),
	}
}

func NewQueries(
	boilerIDB boil.ContextExecutor,
	filesPath string,
) persistence.Queries {
	return persistence.Queries{
		User:        boilerQuery.NewUser(boilerIDB),
		Joke:        boilerQuery.NewJoke(boilerIDB),
		Like:        boilerQuery.NewLike(boilerIDB),
		Comment:     boilerQuery.NewComment(boilerIDB),
		Room:        boilerQuery.NewRoom(boilerIDB),
		RoomMember:  boilerQuery.NewRoomMember(boilerIDB),
		RoomMessage: boilerQuery.NewRoomMessage(boilerIDB),
		FileInfo:    boilerQuery.NewFileInfo(boilerIDB),
		FileContent: localFileSaverQuery.New(filesPath),
	}
}
