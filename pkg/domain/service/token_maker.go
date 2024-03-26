package service

import (
	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
)

var (
	ErrInvalidToken = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided invalid token")
	ErrExpiredToken = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided expired token")
)

type ITokenMaker interface {
	// CreateAccessToken creates access token with given userID
	CreateAccessToken(userID string) (string, error)

	// CreateRefreshToken creates refresh token with given userID
	CreateRefreshToken(userID string) (string, error)

	// VerifyToken verifies given token and, if it's correct, returns its payload
	VerifyToken(token string) (AuthPayload, error)
}

type AuthPayload struct {
	UserID    string
	IsRefresh bool
}
