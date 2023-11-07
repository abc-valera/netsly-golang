package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/spec"
)

var (
	ErrUserNotFound                  = codeerr.NewMsgErr(codeerr.CodeNotFound, "User not found")
	ErrUserWithUsernameAlreadyExists = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "User with such email already exists")

	ErrUsersOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'username', 'fullname' and 'created_at' fields")
)

type IUserRepository interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, userID string, req UserUpdateRequest) error
	Delete(ctx context.Context, userID string) error

	common.Transactioneer
}

func ValidateUserSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "username" && params.OrderBy != "fullname" && params.OrderBy != "created_at" {
		return ErrUsersOrderByNotSupported
	}
	return nil
}

type UserUpdateRequest struct {
	Username string
	Email    string
	Password string
	Fullname string
	Status   string
}

func NewUserUpdateRequest(username, email, password, fullname, status string) (UserUpdateRequest, error) {
	if username == "" {
		return UserUpdateRequest{}, entity.ErrUserUsernameInvalid
	}
	if email == "" {
		return UserUpdateRequest{}, entity.ErrUserEmailInvalid
	}
	if password == "" {
		return UserUpdateRequest{}, entity.ErrUserPasswordInvalid
	}

	return UserUpdateRequest{
		Username: username,
		Email:    email,
		Password: password,
		Fullname: fullname,
		Status:   status,
	}, nil
}
