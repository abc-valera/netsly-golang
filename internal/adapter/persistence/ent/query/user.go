package query

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/err-handler"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
)

type userQuery struct {
	*ent.Client
}

func NewUserQuery(client *ent.Client) query.IUserQuery {
	return &userQuery{
		Client: client,
	}
}

func (uq *userQuery) GetOne(ctx context.Context, fields query.UserGetFields) (*model.User, error) {
	query := uq.User.Query()

	// Where
	if fields.ID != "" {
		query.Where(user.ID(fields.ID))
	}
	if fields.Username != "" {
		query.Where(user.Username(fields.Username))
	}
	if fields.Email != "" {
		query.Where(user.Email(fields.Email))
	}

	entUser, err := query.Only(ctx)
	return dto.FromEntUserToUser(entUser), errhandler.HandleErr(err)
}

func (uq *userQuery) GetAll(ctx context.Context, params query.UserManySelectParams) (model.Users, error) {
	query := uq.User.Query()

	// Where
	if params.SearchBy.Email != "" {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like("email", "%"+params.SearchBy.Email+"%"))
		})
	}
	if params.SearchBy.Username != "" {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like("username", "%"+params.SearchBy.Username+"%"))
		})
	}
	if params.SearchBy.Fullname != "" {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like("fullname", "%"+params.SearchBy.Fullname+"%"))
		})
	}

	// Order
	orderByField := "created_at"
	if params.OrderBy.Username == true {
		orderByField = "username"
	}

	if params.Order == "asc" {
		query.Order(ent.Asc(orderByField))
	} else {
		query.Order(ent.Desc(orderByField))
	}

	// Limit and Offset
	query.Limit(params.Limit)
	query.Offset(params.Offset)

	entUsers, err := query.All(ctx)
	return dto.FromEntUsersToUsers(entUsers), errhandler.HandleErr(err)
}
