package tokenMaker

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/golang-jwt/jwt/v5"
)

type jwtTokenMaker struct {
	accessDuration  time.Duration
	refreshDuration time.Duration

	signMethod jwt.SigningMethod
	signKey    string
}

func NewJWT(
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	signKey string,
) service.ITokenMaker {
	if len(signKey) < 32 {
		coderr.Fatal("JWT_SIGN_KEY environmental variable is invalid")
	}

	return &jwtTokenMaker{
		accessDuration:  accessTokenDuration,
		refreshDuration: refreshTokenDuration,
		signMethod:      jwt.SigningMethodHS256,
		signKey:         signKey,
	}
}

func (s jwtTokenMaker) createToken(userID string, isRefresh bool, duration time.Duration) (string, error) {
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
		return "", coderr.NewInternalErr(err)
	}

	return tokenString, nil
}

func (s jwtTokenMaker) CreateAccessToken(userID string) (string, error) {
	return s.createToken(userID, false, s.accessDuration)
}

func (s jwtTokenMaker) CreateRefreshToken(userID string) (string, error) {
	return s.createToken(userID, true, s.refreshDuration)
}

func (s *jwtTokenMaker) VerifyToken(token string) (service.AuthPayload, error) {
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
