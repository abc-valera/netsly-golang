package auth

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken        = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided invalid token")
	ErrExpiredToken        = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided expired token")
	ErrProvidedAccessToken = coderr.NewCodeMessage(coderr.CodeUnauthenticated, "Provided access token")
)

type Manager struct {
	signKey         string
	signMethod      jwt.SigningMethod
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewManager(
	signKey string,
	accessDuration time.Duration,
	refreshDuration time.Duration,
) Manager {
	if len(signKey) < 32 {
		coderr.Fatal("Sign Key for JWT is invalid")
	}
	return Manager{
		signKey:         signKey,
		signMethod:      jwt.SigningMethodHS256,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

type Payload struct {
	UserID    string
	IsRefresh bool
}

func (m Manager) createToken(payload Payload, duration time.Duration) (string, error) {
	token := jwt.New(m.signMethod)

	claims := jwt.MapClaims{
		"user_id":    payload.UserID,
		"is_refresh": payload.IsRefresh,
		"issued_at":  time.Now().Format(time.RFC3339),
		"expired_at": time.Now().Add(duration).Format(time.RFC3339),
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(m.signKey))
	if err != nil {
		return "", coderr.NewInternalErr(err)
	}

	return tokenString, nil
}

func (m Manager) CreateAccessToken(userID string) (string, error) {
	return m.createToken(
		Payload{
			UserID:    userID,
			IsRefresh: false,
		},
		m.accessDuration,
	)
}

func (m Manager) CreateRefreshToken(userID string) (string, error) {
	return m.createToken(
		Payload{
			UserID:    userID,
			IsRefresh: true,
		},
		m.refreshDuration,
	)
}

func (m Manager) VerifyToken(tokenString string) (Payload, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.signKey), nil
		},
	)
	if err != nil {
		return Payload{}, ErrInvalidToken
	}

	expiredAt, err := time.Parse(time.RFC3339, claims["expired_at"].(string))
	if err != nil {
		return Payload{}, coderr.NewInternalErr(err)
	}

	if time.Now().After(expiredAt) {
		return Payload{}, ErrExpiredToken
	}

	return Payload{
		UserID:    claims["user_id"].(string),
		IsRefresh: claims["is_refresh"].(bool),
	}, nil
}

func (m Manager) RefreshToken(refreshToken string) (string, error) {
	payload, err := m.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	if !payload.IsRefresh {
		return "", ErrProvidedAccessToken
	}

	access, err := m.CreateAccessToken(payload.UserID)
	if err != nil {
		return "", err
	}

	return access, nil
}
