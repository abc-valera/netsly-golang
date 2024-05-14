package service

import "github.com/abc-valera/netsly-api-golang/internal/core/coderr"

var ErrPasswordDontMatch = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Provided password doesn't match the original one")

type IPasswordMaker interface {
	// HashPassword returns hash of the provided password
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash,
	// if matches returns nil, else returns error
	CheckPassword(password, hash string) error
}
