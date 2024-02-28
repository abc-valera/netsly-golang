package service

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrInvalidToken = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid token")
	ErrExpiredToken = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided expired token")
)

type ITokenMaker interface {
	// CreateAccessToken creates access token with given userID
	CreateAccessToken(userID string) (string, AuthPayload, error)

	// CreateRefreshToken creates refresh token with given userID
	CreateRefreshToken(userID string) (string, AuthPayload, error)

	// VerifyToken verifies given token and, if it's correct, returns its payload
	VerifyToken(token string) (AuthPayload, error)
}

// AuthPayload is data which will be stored inside the token
type AuthPayload struct {
	UserID    string
	IsRefresh bool
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewAuthPayload(userID string, isRefresh bool, duration time.Duration) (AuthPayload, error) {
	if userID == "" {
		return AuthPayload{}, coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID for the token")
	}
	if duration == 0 {
		return AuthPayload{}, coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid token duration")
	}

	return AuthPayload{
		UserID:    userID,
		IsRefresh: isRefresh,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (p AuthPayload) Valid() bool {
	return p.ExpiredAt.After(time.Now())
}
