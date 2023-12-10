package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type IUserQuery interface {
	GetAll(ctx context.Context, params UserManySelectParams) (model.Users, error)
	GetOne(ctx context.Context, field UserGetFields) (*model.User, error)
}

type UserSearchByFields struct {
	Username string
	Email    string
	Fullname string
}

type UserOrderByFields struct {
	Username  bool
	CreatedAt bool
}

type UserManySelectParams struct {
	SearchBy UserSearchByFields
	OrderBy  UserOrderByFields
	spec.SelectParams
}

func NewUserOneSelectParams(
	searchBy UserSearchByFields,
) UserManySelectParams {
	return UserManySelectParams{
		SearchBy: searchBy,
	}
}

func NewUserSelectParams(
	searchBy UserSearchByFields,
	orderBy UserOrderByFields,
	order string,
	limit int,
	offset int,
) (UserManySelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return UserManySelectParams{}, err
	}
	return UserManySelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type UserGetFields struct {
	ID       string
	Username string
	Email    string
}
