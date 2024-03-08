package sqlboilerquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/sqlboiler-query/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	mod := sqlboiler.UserWhere.Username.EQ(username)
	return dto.ToDomainUserWithErrHandle(sqlboiler.Users(mod).One(ctx, u.executor))
}

func (u user) GetByEmail(ctx context.Context, email string) (model.User, error) {
	mod := sqlboiler.UserWhere.Email.EQ(email)
	return dto.ToDomainUserWithErrHandle(sqlboiler.Users(mod).One(ctx, u.executor))
}

func (u user) SearchAllByUsername(ctx context.Context, keyword string, params spec.SelectParams) (model.Users, error) {
	mods := common.ToBoilerSelectParamsPipe(
		params,
		qm.Where("username LIKE ?", "%"+keyword+"%"),
	)
	return dto.ToDomainUsersWithErrHandle(sqlboiler.Users(mods...).All(ctx, u.executor))
}
