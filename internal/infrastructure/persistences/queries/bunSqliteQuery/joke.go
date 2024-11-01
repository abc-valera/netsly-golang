package bunSqliteQuery

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
	"github.com/uptrace/bun"
)

type joke struct {
	queryGeneric.IGetOneGetMany[model.Joke]
}

func NewJoke(db bun.IDB) query.IJoke {
	return &joke{
		IGetOneGetMany: bunSqliteQueryGeneric.New(db, bunSqliteDto.NewJoke),
	}
}
