package token

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
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
		global.Log().Fatal("'JWT_SIGN_KEY' environmental variable is invalid")
	}

	return &jwtToken{
		accessDuration:  accessTokenDuration,
		refreshDuration: refreshTokenDuration,
		signMethod:      jwt.SigningMethodHS256,
		signKey:         signKey,
	}
}

func (s jwtToken) createToken(userID string, isRefresh bool, duration time.Duration) (string, service.Payload, error) {
	payload, err := service.NewPayload(userID, isRefresh, duration)
	if err != nil {
		return "", service.Payload{}, err
	}

	token := jwt.New(s.signMethod)

	claims := jwt.MapClaims{
		"user_id":    payload.UserID,
		"is_refresh": payload.IsRefresh,
		"issued_at":  payload.IssuedAt.Format(time.RFC3339),
		"expired_at": payload.ExpiredAt.Format(time.RFC3339),
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		return "", service.Payload{}, coderr.NewInternal(err)
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
		return []byte(s.signKey), nil
	})
	if err != nil {
		return service.Payload{}, service.ErrInvalidToken
	}

	issuedAt, err := time.Parse(time.RFC3339, claims["issued_at"].(string))
	if err != nil {
		return service.Payload{}, coderr.NewInternal(err)
	}
	expiredAt, err := time.Parse(time.RFC3339, claims["expired_at"].(string))
	if err != nil {
		return service.Payload{}, coderr.NewInternal(err)
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
