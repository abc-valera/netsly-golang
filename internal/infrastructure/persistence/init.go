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

	commands := implementation.NewCommands(implementation.CommandsDependencies{
		Boiler: deps.Boiler,
	})

	queries := implementation.NewQueries(implementation.QueriesDependencies{
		Boiler: deps.Boiler,
	})

	transactor := infraCommandTransactor.New(deps)

	return commands, queries, transactor
}
