package auth

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/contexts"
)

type Handler struct {
	authManager Manager
}

func NewHandler(authManager Manager) Handler {
	return Handler{
		authManager: authManager,
	}
}

func (h Handler) HandleBearerAuth(ctx context.Context, operationName string, t ogen.BearerAuth) (context.Context, error) {
	payload, err := h.authManager.VerifyToken(t.Token)
	if err != nil {
		return ctx, err
	}
	return contexts.SetUserID(ctx, payload.UserID), nil
}
