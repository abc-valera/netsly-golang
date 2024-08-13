package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) query.IUser {
	return &user{
		db: db,
	}
}

func (q user) GetByID(ctx context.Context, id string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	res := q.db.Where("id = ?", id).First(&user)
	return gormSqliteDto.NewDomainUser(user), gormSqliteErrutil.HandleQueryResult(res)
}

func (q user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	res := q.db.WithContext(ctx).Where("username = ?", username).First(&user)
	return gormSqliteDto.NewDomainUser(user), gormSqliteErrutil.HandleQueryResult(res)
}

func (q user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	res := q.db.WithContext(ctx).Where("email = ?", email).First(&user)
	return gormSqliteDto.NewDomainUser(user), gormSqliteErrutil.HandleQueryResult(res)
}

func (q user) SearchAllByUsername(ctx context.Context, keyword string, selector selector.Selector) (model.Users, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var users gormSqliteDto.Users
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("username LIKE ?", "%"+keyword+"%").
		Find(&users)
	return gormSqliteDto.NewDomainUsers(users), gormSqliteErrutil.HandleQueryResult(res)
}
