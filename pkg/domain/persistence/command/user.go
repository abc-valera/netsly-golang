package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/core/optional"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
)

var (
	ErrUserWithUsernameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such email already exists")
)

type IUser interface {
	Create(ctx context.Context, req model.User) (model.User, error)
	Update(ctx context.Context, id string, req UserUpdate) (model.User, error)
	Delete(ctx context.Context, id string) error
}

type UserUpdate struct {
	HashedPassword optional.Optional[string]
	Fullname       optional.Optional[string]
	Status         optional.Optional[string]
}
