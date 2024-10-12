package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrors"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/gormSqliteQuery/gormSelector"
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
	return user.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}

func (q user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	res := q.db.WithContext(ctx).Where("username = ?", username).First(&user)
	return user.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}

func (q user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	res := q.db.WithContext(ctx).Where("email = ?", email).First(&user)
	return user.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}

func (q user) GetAll(ctx context.Context, s selector.Selector) (model.Users, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var users gormSqliteDto.Users
	res := gormSelector.WithSelector(q.db, s).Find(&users)
	return users.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}
