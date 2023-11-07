package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity/common"
)

var (
	ErrUserUsernameInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid username")
	ErrUserEmailInvalid    = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid email")
	ErrUserPasswordInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid hashed password")
	ErrUserFullnameInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid fullname")
	ErrUserStatusInvalid   = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid status")
)

type User struct {
	common.BaseEntity
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
}

func NewUser(username, email, hashedPassword, fullname, status string) (*User, error) {
	if username == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid username")
	}
	if email == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid email")
	}
	if hashedPassword == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid hashed password")
	}

	return &User{
		BaseEntity:     common.NewBaseEntity(),
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Fullname:       fullname,
		Status:         status,
	}, nil
}

type Users []*User
