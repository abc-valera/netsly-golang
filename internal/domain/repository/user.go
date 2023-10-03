package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

var (
	ErrUserNotFound                  = codeerr.NewErrWithMsg(codeerr.CodeNotFound, "User not found")
	ErrUserWithUsernameAlreadyExists = codeerr.NewErrWithMsg(codeerr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = codeerr.NewErrWithMsg(codeerr.CodeAlreadyExists, "User with such email already exists")
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error

	Transactioneer
}
