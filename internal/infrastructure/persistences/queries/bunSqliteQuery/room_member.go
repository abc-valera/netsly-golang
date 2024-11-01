package bunSqliteQuery

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
	"github.com/uptrace/bun"
)

type roomMember struct {
	queryGeneric.IGetOneGetMany[model.RoomMember]
}

func NewRoomMember(db bun.IDB) query.IRoomMember {
	return &roomMember{
		IGetOneGetMany: bunSqliteQueryGeneric.New(db, bunSqliteDto.NewRoomMember),
	}
}
