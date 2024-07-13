package persistence

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	domainCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	infraCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
)

func New(postgresUrl string) (
	persistence.Commands,
	persistence.Queries,
	domainCommandTransactor.ITransactor,
) {
	deps := implementation.NewPersistenceDependencies(postgresUrl)

	commands := implementation.NewCommands(deps.BoilerDB)

	queries := implementation.NewQueries(deps.BoilerDB)

	transactor := infraCommandTransactor.New(deps)

	return commands, queries, transactor
}
