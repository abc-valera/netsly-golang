package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

var (
	ErrUserWithUsernameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such email already exists")
)

type IUser interface {
	Create(ctx context.Context, req model.User) (model.User, error)
	Update(ctx context.Context, id string, req UserUpdateRequest) (model.User, error)
	Delete(ctx context.Context, id string) error
}

type UserUpdateRequest struct {
	HashedPassword *string
	Fullname       *string
	Status         *string
}
