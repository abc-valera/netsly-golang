package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/google/uuid"
)

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
	CreatedAt      time.Time
}

func NewUser(username, email, hashedPassword, fullname, status string) (*User, error) {
	if username == "" {
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid username")
	}
	if email == "" {
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid email")
	}
	if hashedPassword == "" {
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid hashed password")
	}

	return &User{
		ID:             uuid.NewString(),
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Fullname:       fullname,
		Status:         status,
		CreatedAt:      time.Now(),
	}, nil
}

type Users []*User
