package service

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
)

var (
	ErrInvalidToken = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid token")
	ErrExpiredToken = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided expired token")
)

type ITokenMaker interface {
	// CreateAccessToken creates access token with given userID
	CreateAccessToken(userID string) (string, Payload, error)

	// CreateRefreshToken creates refresh token with given userID
	CreateRefreshToken(userID string) (string, Payload, error)

	// VerifyToken verifies given token and, if it's correct, returns its payload
	VerifyToken(token string) (Payload, error)
}

// Payload is data which will be stored inside the token
type Payload struct {
	UserID    string
	IsRefresh bool
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewPayload(userID string, isRefresh bool, duration time.Duration) (Payload, error) {
	if userID == "" {
		return Payload{}, coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID for the token")
	}
	if duration == 0 {
		return Payload{}, coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid token duration")
	}

	return Payload{
		UserID:    userID,
		IsRefresh: isRefresh,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (p Payload) Valid() bool {
	return p.ExpiredAt.After(time.Now())
}
