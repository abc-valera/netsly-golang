package bunSqliteQuery

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
	"github.com/uptrace/bun"
)

type user struct {
	queryGeneric.IGetOneGetMany[model.User]
}

func NewUser(db bun.IDB) query.IUser {
	return &user{
		IGetOneGetMany: bunSqliteQueryGeneric.New(db, bunSqliteDto.NewUser),
	}
}
