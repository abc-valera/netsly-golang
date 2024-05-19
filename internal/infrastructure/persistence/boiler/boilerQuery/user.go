package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerQuery/selector"
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

func (u user) GetByID(ctx context.Context, id string) (model.User, error) {
	return boilerDto.NewDomainUserWithErrHandle(sqlboiler.FindUser(ctx, u.executor, id))
}

func (u user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	query := sqlboiler.UserWhere.Username.EQ(username)
	return boilerDto.NewDomainUserWithErrHandle(sqlboiler.Users(query).One(ctx, u.executor))
}

func (u user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	query := sqlboiler.UserWhere.Email.EQ(email)
	return boilerDto.NewDomainUserWithErrHandle(sqlboiler.Users(query).One(ctx, u.executor))
}

func (u user) SearchAllByUsername(ctx context.Context, keyword string, params selector1.Selector) (model.Users, error) {
	queries := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.UserWhere.Username.LIKE("%"+keyword+"%"),
	)
	return boilerDto.NewDomainUsersWithErrHandle(sqlboiler.Users(queries...).All(ctx, u.executor))
}
