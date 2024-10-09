package session

import (
	"context"
	"net/http"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
)

func SetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), UserIDKey, userID))
}

func GetUserID(r *http.Request) string {
	// TODO: do this in a separate function: webAppContexts.GetUserID
	id, ok := r.Context().Value(UserIDKey).(string)
	if !ok {
		global.Log().Error("failed to get user id from context")
		return ""
	}
	return id
}
