package sqlboilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/sqlboilerQuery/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type user struct {
	executor boil.ContextExecutor
}

func newUser(executor boil.ContextExecutor) query.IUser {
	return &user{
		executor: executor,
	}
}

func (u user) GetByID(ctx context.Context, id string) (model.User, error) {
	return dto.ToDomainUserWithErrHandle(sqlboiler.FindUser(ctx, u.executor, id))
}

func (u user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	query := sqlboiler.UserWhere.Username.EQ(username)
	return dto.ToDomainUserWithErrHandle(sqlboiler.Users(query).One(ctx, u.executor))
}

func (u user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	query := sqlboiler.UserWhere.Email.EQ(email)
	return dto.ToDomainUserWithErrHandle(sqlboiler.Users(query).One(ctx, u.executor))
}

func (u user) SearchAllByUsername(ctx context.Context, keyword string, params spec.SelectParams) (model.Users, error) {
	queries := common.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.UserWhere.Username.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainUsersWithErrHandle(sqlboiler.Users(queries...).All(ctx, u.executor))
}
