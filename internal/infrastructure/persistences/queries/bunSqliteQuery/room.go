package bunSqliteQuery

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
	"github.com/uptrace/bun"
)

type room struct {
	queryGeneric.IGetOneGetMany[model.Room]
}

func NewRoom(db bun.IDB) query.IRoom {
	return &room{
		IGetOneGetMany: bunSqliteQueryGeneric.New(db, bunSqliteDto.NewRoom),
	}
}
