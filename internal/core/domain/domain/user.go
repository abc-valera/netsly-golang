package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

var (
	ErrUserIDInvalid       = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	ErrUserUsernameInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid username")
	ErrUserEmailInvalid    = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid email")
	ErrUserPasswordInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid hashed password")
	ErrUserFullnameInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid fullname")
	ErrUserStatusInvalid   = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid status")
)

// UserDomain is responsible for validation and handling user domain logic
type UserDomain struct {
	// Data layer access
	query   query.IUserQuery
	command command.IUserCommand

	// Service layer access
	passMaker service.IPasswordMaker
}

func NewUserDomain(
	query query.IUserQuery,
	command command.IUserCommand,
	passMaker service.IPasswordMaker,
) UserDomain {
	return UserDomain{
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

func (u UserDomain) Create(ctx context.Context, req UserCreateRequest) error {
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

func (u UserDomain) Update(ctx context.Context, userID string, req UserUpdateRequest) error {
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
	return u.command.Update(ctx, userID, model.UserUpdate{
		HashedPassword: &hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
	})
}

type UserDeleteRequest struct {
	Password string
}

func (u UserDomain) Delete(ctx context.Context, userID string, req UserDeleteRequest) error {
	// Validation
	if userID == "" {
		return ErrUserIDInvalid
	}

	// Domain logic
	user, err := u.query.GetOne(ctx, query.NewUserOneSelectParams(
		query.UserSearchByFields{ID: userID},
	))
	if err != nil {
		return err
	}

	if err := u.passMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	// Delete from data source
	return u.command.Delete(ctx, userID)
}
