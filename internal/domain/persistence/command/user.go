package command

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var (
	ErrUserWithUsernameAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "User with such email already exists")
)

type IUser interface {
	Create(ctx context.Context, req model.User) (model.User, error)
	Update(ctx context.Context, ids model.User, req UserUpdateRequest) (model.User, error)
	Delete(ctx context.Context, req model.User) error
}

type UserUpdateRequest struct {
	UpdatedAt time.Time

	HashedPassword *string
	Fullname       *string
	Status         *string
}
