package token

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "12345678901234567890123456789012"

type jwtToken struct {
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewTokenMaker(accessDuration, refreshDuration time.Duration) service.ITokenMaker {
	return &jwtToken{
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (s jwtToken) createToken(userID string, isRefresh bool, duration time.Duration) (string, service.Payload, error) {
	payload, err := service.NewPayload(userID, isRefresh, duration)
	if err != nil {
		return "", service.Payload{}, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{
		"user_id":    payload.UserID,
		"is_refresh": payload.IsRefresh,
		"issued_at":  payload.IssuedAt.Format(time.RFC3339),
		"expired_at": payload.ExpiredAt.Format(time.RFC3339),
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", service.Payload{}, codeerr.NewInternal("jwtToken.createToken", err)
	}

	return tokenString, payload, nil
}

func (s jwtToken) CreateAccessToken(userID string) (string, service.Payload, error) {
	return s.createToken(userID, false, s.accessDuration)
}

func (s jwtToken) CreateRefreshToken(userID string) (string, service.Payload, error) {
	return s.createToken(userID, true, s.refreshDuration)
}

func (s *jwtToken) VerifyToken(token string) (service.Payload, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return service.Payload{}, service.ErrInvalidToken
	}

	issuedAt, err := time.Parse(time.RFC3339, claims["issued_at"].(string))
	if err != nil {
		return service.Payload{}, codeerr.NewInternal("jwtToken.VerifyToken", err)
	}
	expiredAt, err := time.Parse(time.RFC3339, claims["expired_at"].(string))
	if err != nil {
		return service.Payload{}, codeerr.NewInternal("jwtToken.VerifyToken", err)
	}

	var payload service.Payload
	payload.UserID = claims["user_id"].(string)
	payload.IsRefresh = claims["is_refresh"].(bool)
	payload.IssuedAt = issuedAt
	payload.ExpiredAt = expiredAt

	if ok := payload.Valid(); !ok {
		return service.Payload{}, service.ErrExpiredToken
	}

	return payload, nil
}
