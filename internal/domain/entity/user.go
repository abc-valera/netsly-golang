package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

// User is responsible for validation and handling user domain logic
type User struct {
	// Data layer access
	query   query.IUser
	command command.IUser

	// Service layer access
	passMaker service.IPasswordMaker
}

func NewUser(
	command command.IUser,
	query query.IUser,
	passMaker service.IPasswordMaker,
) User {
	return User{
		query:     query,
		command:   command,
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

func (u User) Create(ctx context.Context, req UserCreateRequest) (model.User, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.User{}, err
	}

	baseModel := common.NewBaseEntity()

	hashedPassword, err := u.passMaker.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	return u.command.Create(ctx, model.User{
		BaseEntity:     baseModel,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserUpdateRequest struct {
	Password *string `validate:"min=2,max=32"`
	Fullname *string `validate:"max=64"`
	Status   *string `validate:"max=128"`
}

func (u User) Update(ctx context.Context, userID string, req UserUpdateRequest) (model.User, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.User{}, err
	}

	// Domain logic
	hashedPassword, err := u.passMaker.HashPassword(*req.Password)
	if err != nil {
		return model.User{}, err
	}

	// Edit in data source
	return u.command.Update(ctx, userID, command.UserUpdate{
		HashedPassword: &hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserDeleteRequest struct {
	Password string `validate:"required,min=2,max=32"`
}

func (u User) Delete(ctx context.Context, userID string, req UserDeleteRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	// Domain logic
	user, err := u.query.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := u.passMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	// Delete from data source
	return u.command.Delete(ctx, userID)
}
