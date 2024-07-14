package persistence

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	domainCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	infraCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
)

func New(
	postgresUrl string,
	filesPath string,
) (
	persistence.Commands,
	persistence.Queries,
	domainCommandTransactor.ITransactor,
) {
	deps := implementation.NewPersistenceDependencies(
		postgresUrl,
		filesPath,
	)

	commands := implementation.NewCommands(
		deps.BoilerDB,
		deps.FilesPath,
	)

	queries := implementation.NewQueries(
		deps.BoilerDB,
		deps.FilesPath,
	)

	transactor := infraCommandTransactor.New(deps)

	return commands, queries, transactor
}
