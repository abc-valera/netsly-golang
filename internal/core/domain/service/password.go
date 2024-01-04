package service

import "github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"

var (
	ErrInvalidPassword = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided password is invalid")
)

type IPasswordMaker interface {
	// HashPassword returns hash of the provided password
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash,
	// if matches returns nil, else returns error
	CheckPassword(password, hash string) error
}
