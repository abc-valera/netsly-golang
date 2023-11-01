package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
)

var (
	ErrUserNotFound                  = codeerr.NewMsgErr(codeerr.CodeNotFound, "User not found")
	ErrUserWithUsernameAlreadyExists = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "User with such email already exists")
)

type IUserRepository interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userID string) error

	common.Transactioneer
}
