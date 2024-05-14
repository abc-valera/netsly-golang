package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/optional"
	"github.com/abc-valera/netsly-api-golang/internal/core/validator"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/google/uuid"
)

type IUser interface {
	Create(ctx context.Context, req UserCreateRequest) (model.User, error)
	Update(ctx context.Context, userID string, req UserUpdateRequest) (model.User, error)
	Delete(ctx context.Context, userID string, req UserDeleteRequest) error

	query.IUser
}

type user struct {
	command command.IUser
	query.IUser

	validator validator.IValidator
	passMaker service.IPasswordMaker
}

func NewUser(
	command command.IUser,
	query query.IUser,
	validator validator.IValidator,
	passMaker service.IPasswordMaker,
) IUser {
	return user{
		command:   command,
		IUser:     query,
		validator: validator,
		passMaker: passMaker,
	}
}

type UserCreateRequest struct {
	Username string `validate:"required,min=2,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=2,max=32"`
	Fullname string `validate:"max=64"`
	Status   string `validate:"max=128"`
}

func (u user) Create(ctx context.Context, req UserCreateRequest) (model.User, error) {
	if err := u.validator.Struct(req); err != nil {
		return model.User{}, err
	}

	hashedPassword, err := u.passMaker.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	return u.command.Create(ctx, model.User{
		ID:             uuid.New().String(),
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		CreatedAt:      time.Now(),
	})
}

type UserUpdateRequest struct {
	Password optional.Optional[string] `validate:"min=2,max=32"`
	Fullname optional.Optional[string] `validate:"max=64"`
	Status   optional.Optional[string] `validate:"max=128"`
}

func (u user) Update(ctx context.Context, userID string, req UserUpdateRequest) (model.User, error) {
	if err := u.validator.Struct(req); err != nil {
		return model.User{}, err
	}

	hashedPassword, err := u.passMaker.HashPassword(req.Password.Value())
	if err != nil {
		return model.User{}, err
	}

	return u.command.Update(ctx, userID, command.UserUpdate{
		HashedPassword: optional.NewOptional(hashedPassword),
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserDeleteRequest struct {
	Password string
}

func (u user) Delete(ctx context.Context, userID string, req UserDeleteRequest) error {
	if err := u.validator.Struct(req); err != nil {
		return err
	}

	user, err := u.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := u.passMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	return u.command.Delete(ctx, userID)
}
