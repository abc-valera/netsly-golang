package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/dto"
)

type userQuery struct {
	*ent.Client
}

func NewUserQuery(client *ent.Client) query.IUser {
	return &userQuery{
		Client: client,
	}
}

func (uq userQuery) GetByID(ctx context.Context, id string) (model.User, error) {
	return dto.FromEntUserWithErrHandle(uq.User.Get(ctx, id))
}

func (uq userQuery) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return dto.FromEntUserWithErrHandle(
		uq.User.Query().
			Where(user.Username(username)).
			Only(ctx),
	)
}

func (uq userQuery) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return dto.FromEntUserWithErrHandle(
		uq.User.
			Query().
			Where(user.Email(email)).
			Only(ctx),
	)
}

func (uq userQuery) SearchAllByUsername(ctx context.Context, keyword string, params spec.SelectParams) (model.Users, error) {
	query := uq.User.
		Query().
		Where(user.UsernameContains(keyword))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntUsersWithErrHandle(query.All(ctx))
}
