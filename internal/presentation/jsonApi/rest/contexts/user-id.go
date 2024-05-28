package contexts

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

type key string

const (
	userIDKey key = "user-id"
)

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok {
		return "", coderr.NewInternalString("UserID not found in context")
	}
	return userID, nil
}
