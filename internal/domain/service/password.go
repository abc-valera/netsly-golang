package service

import "github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"

var (
	ErrInvalidPassword = codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided password is invalid")
)

type PasswordMaker interface {
	// HashPassword returns hash of the provided password
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash,
	// if matches returns nil, else returns error
	CheckPassword(password, hashedPassword string) error
}
