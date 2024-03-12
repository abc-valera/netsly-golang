package tokenMaker

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/golang-jwt/jwt/v5"
)

type jwtToken struct {
	accessDuration  time.Duration
	refreshDuration time.Duration

	signMethod jwt.SigningMethod
	signKey    string
}

func NewTokenMaker(
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	signKey string,
) service.ITokenMaker {
	if len(signKey) < 32 {
		coderr.Fatal("JWT_SIGN_KEY environmental variable is invalid")
	}

	return &jwtToken{
		accessDuration:  accessTokenDuration,
		refreshDuration: refreshTokenDuration,
		signMethod:      jwt.SigningMethodHS256,
		signKey:         signKey,
	}
}

func (s jwtToken) createToken(userID string, isRefresh bool, duration time.Duration) (string, service.AuthPayload, error) {
	payload := service.AuthPayload{
		UserID:    userID,
		IsRefresh: isRefresh,
	}

	token := jwt.New(s.signMethod)

	claims := jwt.MapClaims{
		"user_id":    payload.UserID,
		"is_refresh": payload.IsRefresh,
		"issued_at":  time.Now().Format(time.RFC3339),
		"expired_at": time.Now().Add(duration).Format(time.RFC3339),
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		return "", service.AuthPayload{}, coderr.NewInternalErr(err)
	}

	return tokenString, payload, nil
}

func (s jwtToken) CreateAccessToken(userID string) (string, service.AuthPayload, error) {
	return s.createToken(userID, false, s.accessDuration)
}

func (s jwtToken) CreateRefreshToken(userID string) (string, service.AuthPayload, error) {
	return s.createToken(userID, true, s.refreshDuration)
}

func (s *jwtToken) VerifyToken(token string) (service.AuthPayload, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
	})
	if err != nil {
		return service.AuthPayload{}, service.ErrInvalidToken
	}

	expiredAt, err := time.Parse(time.RFC3339, claims["expired_at"].(string))
	if err != nil {
		return service.AuthPayload{}, coderr.NewInternalErr(err)
	}

	if time.Now().After(expiredAt) {
		return service.AuthPayload{}, service.ErrExpiredToken
	}

	return service.AuthPayload{
		UserID:    claims["user_id"].(string),
		IsRefresh: claims["is_refresh"].(bool),
	}, nil
}
