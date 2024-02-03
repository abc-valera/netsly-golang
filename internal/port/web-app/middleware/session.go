package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/core/service"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/cookie"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
)

func NewSessionMiddleware(tokenMaker service.ITokenMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				accessToken, err := cookie.Get(r, cookie.AccessTokenKey)
				if err != nil && err != cookie.ErrNoCookie && err != cookie.ErrInvalidValue {
					global.Log.Error("failed to get access token", "err", err)
				}

				if accessToken != "" {
					payload, err := tokenMaker.VerifyToken(accessToken)
					if err != nil && err != service.ErrExpiredToken && err != service.ErrInvalidToken {
						global.Log.Error("failed to verify access token", "err", err)
					} else {
						if strings.HasPrefix(r.URL.Path, "/sign") {
							http.Redirect(w, r, "/home", http.StatusMovedPermanently)
						}
						next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserIDKey, payload.UserID)))
						return
					}
				}

				refreshToken, err := cookie.Get(r, cookie.RefreshTokenKey)
				if err != nil && err != cookie.ErrNoCookie && err != cookie.ErrInvalidValue {
					global.Log.Error("failed to get refresh token", "err", err)
				}

				if refreshToken != "" {
					payload, err := tokenMaker.VerifyToken(refreshToken)
					if err != nil && err != service.ErrExpiredToken && err != service.ErrInvalidToken {
						global.Log.Error("failed to verify access token", "err", err)
					} else {
						if strings.HasPrefix(r.URL.Path, "/sign") {
							http.Redirect(w, r, "/home", http.StatusMovedPermanently)
						}
						next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserIDKey, payload.UserID)))
						access, _, err := tokenMaker.CreateAccessToken(payload.UserID)
						if err != nil {
							global.Log.Error("failed to create access token", "err", err)
						}
						cookie.Set(w, cookie.AccessTokenKey, access)
						return
					}
				}

				// if the request is not for sign, then redirect the user to the sign page
				if !strings.HasPrefix(r.URL.Path, "/sign") {
					http.Redirect(w, r, "/sign", http.StatusMovedPermanently)
				} else {
					next.ServeHTTP(w, r)
				}
			},
		)
	}
}
