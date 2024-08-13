package infrastructure

import (
	"os"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/commandTransactor"
	"github.com/abc-valera/netsly-golang/internal/core/entityTransactor"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/commandsAndQueries"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/localFileSaver"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/transactors"
)

func NewPersistences(services domain.Services) (
	persistence.Commands,
	persistence.Queries,
	commandTransactor.ITransactor,
	entityTransactor.ITransactor,
) {
	var commandsAndQueriesDependencies commandsAndQueries.Dependencies
	var transactorsDependencies transactors.Dependencies

	if gormSqliteEnv := strings.TrimSpace(os.Getenv("GORM_SQLITE_FOLDER_PATH")); gormSqliteEnv != "" {
		gormSqliteDependency := coderr.Must(gormSqlite.New(gormSqliteEnv))

		commandsAndQueriesDependencies.GormSqlite = gormSqliteDependency
		transactorsDependencies.GormSqlite = gormSqliteDependency
	}

	if boilerSqliteEnv := strings.TrimSpace(os.Getenv("BOILER_SQLITE_FOLDER_PATH")); boilerSqliteEnv != "" {
		boilerSqliteDependency := coderr.Must(boilerSqlite.New(boilerSqliteEnv))

		commandsAndQueriesDependencies.BoilerSqlite = boilerSqliteDependency
		transactorsDependencies.BoilerSqlite = boilerSqliteDependency
	}

	if localFileSaverEnv := strings.TrimSpace(os.Getenv("LOCAL_FILE_SAVER_FOLDER_PATH")); localFileSaverEnv != "" {
		localFileSaverDependency := coderr.Must(localFileSaver.New(localFileSaverEnv))

		commandsAndQueriesDependencies.LocalFileSaver = localFileSaverDependency
		transactorsDependencies.LocalFileSaver = localFileSaverDependency
	}

	commands, queries, err := commandsAndQueries.New(commandsAndQueriesDependencies)
	if err != nil {
		coderr.Fatal(err)
	}

	// Init command transactor
	commandTransactor := transactors.NewCommandTransactor(transactorsDependencies)

	// Init Entity Transactor
	entityTransactor := transactors.NewEntityTransactor(transactorsDependencies, services)

	return commands, queries, commandTransactor, entityTransactor
}
