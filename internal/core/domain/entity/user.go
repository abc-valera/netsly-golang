package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/spec"
)

var (
	ErrUsersOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'username', 'fullname' and 'created_at' fields")
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

func ValidateUserSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "username" && params.OrderBy != "fullname" && params.OrderBy != "created_at" {
		return ErrUsersOrderByNotSupported
	}
	return nil
}
