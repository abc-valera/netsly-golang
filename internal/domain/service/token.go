package service

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrInvalidToken = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided invalid token")
	ErrExpiredToken = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided expired token")
)

type ITokenMaker interface {
	// CreateAccessToken creates access token with given userID
	CreateAccessToken(userID string) (string, AuthPayload, error)

	// CreateRefreshToken creates refresh token with given userID
	CreateRefreshToken(userID string) (string, AuthPayload, error)

	// VerifyToken verifies given token and, if it's correct, returns its payload
	VerifyToken(token string) (AuthPayload, error)
}

type AuthPayload struct {
	UserID    string
	IsRefresh bool
}
