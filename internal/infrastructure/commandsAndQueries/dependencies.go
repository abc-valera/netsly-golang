package commandsAndQueries

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

type Dependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   boil.ContextExecutor
	LocalFileSaver string
}

func New(deps Dependencies) (
	persistence.Commands,
	persistence.Queries,
	error,
) {
	commands, err := newCommands(deps)
	if err != nil {
		return persistence.Commands{}, persistence.Queries{}, err
	}
	queries, err := newQueries(deps)
	if err != nil {
		return persistence.Commands{}, persistence.Queries{}, err
	}
	return commands, queries, nil
}
