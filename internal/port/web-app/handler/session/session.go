package session

import (
	"context"
	"net/http"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
)

func SetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), UserIDKey, userID))
}

func GetUserID(r *http.Request) string {
	return r.Context().Value(UserIDKey).(string)
}
