package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

var (
	ErrUserWithUsernameAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = coderr.NewMessage(coderr.CodeAlreadyExists, "User with such email already exists")
)

type IUser interface {
	Create(ctx context.Context, req model.User) error
	Update(ctx context.Context, id string, req UserUpdate) error
	Delete(ctx context.Context, id string) error
}

type UserUpdate struct {
	HashedPassword *string
	Fullname       *string
	Status         *string
}
