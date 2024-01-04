package query

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type userQuery struct {
	*ent.Client
}

func NewUserQuery(client *ent.Client) query.IUserQuery {
	return &userQuery{
		Client: client,
	}
}

func (uq *userQuery) GetByID(ctx context.Context, id string) (*model.User, error) {
	return dto.FromEntUserToUserWithErrHandle(uq.User.Get(ctx, id))
}

func (uq *userQuery) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return dto.FromEntUserToUserWithErrHandle(
		uq.User.Query().
			Where(user.Username(username)).
			Only(ctx),
	)
}

func (uq *userQuery) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return dto.FromEntUserToUserWithErrHandle(
		uq.User.
			Query().
			Where(user.Email(email)).
			Only(ctx),
	)
}

func (uq *userQuery) SearchAllByUsername(ctx context.Context, keyword string, params spec.SelectParams) (model.Users, error) {
	query := uq.User.
		Query().
		Where(func(s *sql.Selector) { s.Where(sql.Like("username", "%"+keyword+"%")) })

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntUsersToUsersWithErrHandle(query.All(ctx))
}
