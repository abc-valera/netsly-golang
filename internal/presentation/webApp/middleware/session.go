package middleware

import (
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/cookie"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := cookie.Get(r, cookie.UserIDKey)
		if err == nil && userID != "" {
			next.ServeHTTP(w, r)
			return
		}

		if err != nil && err != cookie.ErrInvalidValue && err != cookie.ErrNoCookie {
			global.Log().Error("failed to get user id cookie", "err", err)
			return
		}

		if err == cookie.ErrInvalidValue || err == cookie.ErrNoCookie {
			// if the request is not for sign, then redirect the user to the sign page
			if !strings.HasPrefix(r.URL.Path, "/sign") {
				http.Redirect(w, r, "/sign", http.StatusMovedPermanently)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		if userID == "" {
			global.Log().Error("user id is empty")
			return
		}
	})
}
