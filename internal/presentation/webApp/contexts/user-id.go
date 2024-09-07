package contexts

import (
	"context"
	"net/http"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

type key string

const (
	userIDKey key = "user-id"
)

func SetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userIDKey, userID))
}

func GetUserID(r *http.Request) (string, error) {
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		return "", coderr.NewInternalString("UserID not found in context")
	}
	return userID, nil
}
