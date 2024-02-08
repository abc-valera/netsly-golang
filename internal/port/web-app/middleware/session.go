package middleware

import (
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/cookie"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler/session"
)

func NewSessionMiddleware(tokenMaker service.ITokenMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				accessToken, err := cookie.Get(r, cookie.AccessTokenKey)
				if err != nil && err != cookie.ErrNoCookie && err != cookie.ErrInvalidValue {
					global.Log().Error("failed to get access token", "err", err)
				}

				if accessToken != "" {
					payload, err := tokenMaker.VerifyToken(accessToken)
					if err == nil {
						if strings.HasPrefix(r.URL.Path, "/sign") {
							http.Redirect(w, r, "/home", http.StatusMovedPermanently)
						}
						next.ServeHTTP(w, session.SetUserID(r, payload.UserID))
						return
					}
					if err != service.ErrExpiredToken && err != service.ErrInvalidToken {
						global.Log().Error("failed to verify access token", "err", err)
					}
				}

				refreshToken, err := cookie.Get(r, cookie.RefreshTokenKey)
				if err != nil && err != cookie.ErrNoCookie && err != cookie.ErrInvalidValue {
					global.Log().Error("failed to get refresh token", "err", err)
				}

				if refreshToken != "" {
					payload, err := tokenMaker.VerifyToken(refreshToken)
					if err == nil {
						access, _, err := tokenMaker.CreateAccessToken(payload.UserID)
						if err != nil {
							global.Log().Error("failed to create access token", "err", err)
						}
						cookie.Set(w, cookie.AccessTokenKey, access)
						if strings.HasPrefix(r.URL.Path, "/sign") {
							http.Redirect(w, r, "/home", http.StatusMovedPermanently)
						} else {
							next.ServeHTTP(w, session.SetUserID(r, payload.UserID))
						}
						return
					}
					if err != service.ErrExpiredToken && err != service.ErrInvalidToken {
						global.Log().Error("failed to verify refresh token", "err", err)
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
