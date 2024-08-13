package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteQuery/selector"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type user struct {
	executor boil.ContextExecutor
}

func NewUser(executor boil.ContextExecutor) query.IUser {
	return &user{
		executor: executor,
	}
}

func (q user) GetByID(ctx context.Context, id string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	user, err := sqlboiler.FindUser(ctx, q.executor, id)
	return boilerSqliteDto.NewDomainUser(user), boilerSqliteErrutil.HandleErr(err)
}

func (q user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	query := sqlboiler.UserWhere.Username.EQ(username)
	user, err := sqlboiler.Users(query).One(ctx, q.executor)
	return boilerSqliteDto.NewDomainUser(user), boilerSqliteErrutil.HandleErr(err)
}

func (q user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	query := sqlboiler.UserWhere.Email.EQ(email)
	user, err := sqlboiler.Users(query).One(ctx, q.executor)
	return boilerSqliteDto.NewDomainUser(user), boilerSqliteErrutil.HandleErr(err)
}

func (q user) SearchAllByUsername(ctx context.Context, keyword string, params selector1.Selector) (model.Users, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	queries := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.UserWhere.Username.LIKE("%"+keyword+"%"),
	)
	users, err := sqlboiler.Users(queries...).All(ctx, q.executor)
	return boilerSqliteDto.NewDomainUsers(users), boilerSqliteErrutil.HandleErr(err)
}
