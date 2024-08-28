package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type user struct {
	db bun.IDB
}

func NewUser(db bun.IDB) query.IUser {
	return &user{
		db: db,
	}
}

func (q user) GetByID(ctx context.Context, id string) (model.User, error) {
	bunUser := bunSqliteDto.User{}
	err := q.db.NewSelect().Model(&bunUser).Where("id = ?", id).Scan(ctx)
	return bunUser.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	bunUser := bunSqliteDto.User{}
	err := q.db.NewSelect().Model(&bunUser).Where("username = ?", username).Scan(ctx)
	return bunUser.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	bunUser := bunSqliteDto.User{}
	err := q.db.NewSelect().Model(&bunUser).Where("email = ?", email).Scan(ctx)
	return bunUser.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q user) SearchAllByUsername(ctx context.Context, keyword string, selector selector.Selector) (model.Users, error) {
	bunUsers := bunSqliteDto.Users{}
	err := q.db.NewSelect().Model(&bunUsers).Where("username LIKE ?", "%"+keyword+"%").Scan(ctx)
	return bunUsers.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}
