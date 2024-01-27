package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/service"
)

var (
	ErrUserIDInvalid       = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID")
	ErrUserUsernameInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid username")
	ErrUserEmailInvalid    = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid email")
	ErrUserPasswordInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid hashed password")
	ErrUserFullnameInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid fullname")
	ErrUserStatusInvalid   = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid status")
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
	Username string
	Email    string
	Password string
	Fullname string
	Status   string
}

func (u User) Create(ctx context.Context, req UserCreateRequest) error {
	// Validation
	if req.Username == "" || len(req.Username) < 4 || len(req.Username) > 32 {
		return ErrUserUsernameInvalid
	}
	if req.Email == "" {
		return ErrUserEmailInvalid
	}
	if req.Password == "" || len(req.Password) < 4 || len(req.Password) > 32 {
		return ErrUserPasswordInvalid
	}
	if len(req.Fullname) > 64 {
		return ErrUserFullnameInvalid
	}
	if len(req.Status) > 128 {
		return ErrUserStatusInvalid
	}

	// Domain logic
	baseModel := common.NewBaseModel()

	hashedPassword, err := u.passMaker.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// Save to data source
	return u.command.Create(ctx, model.User{
		BaseModel:      baseModel,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserUpdateRequest struct {
	Password *string
	Fullname *string
	Status   *string
}

func (u User) Update(ctx context.Context, userID string, req UserUpdateRequest) error {
	// Validation
	if userID == "" {
		return ErrUserIDInvalid
	}
	if req.Password != nil || len(*req.Password) < 4 || len(*req.Password) > 32 {
		return ErrUserPasswordInvalid
	}
	if req.Fullname != nil || len(*req.Fullname) > 64 {
		return ErrUserFullnameInvalid
	}
	if req.Status != nil || len(*req.Status) > 128 {
		return ErrUserStatusInvalid
	}

	// Domain logic
	hashedPassword, err := u.passMaker.HashPassword(*req.Password)
	if err != nil {
		return err
	}

	// Edit in data source
	return u.command.Update(ctx, userID, command.UserUpdate{
		HashedPassword: &hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserDeleteRequest struct {
	Password string
}

func (u User) Delete(ctx context.Context, userID string, req UserDeleteRequest) error {
	// Validation
	if userID == "" {
		return ErrUserIDInvalid
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
