package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type Key string

const (
	PayloadKey Key = "auth_payload"
)

func payloadUserID(ctx context.Context) string {
	return ctx.Value(PayloadKey).(service.Payload).UserID
}

type SecurityHandler struct {
	tokenMaker service.ITokenMaker
}

func NewSecurityHandler(tokenMaker service.ITokenMaker) SecurityHandler {
	return SecurityHandler{
		tokenMaker: tokenMaker,
	}
}

func (h SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t ogen.BearerAuth) (context.Context, error) {
	payload, err := h.tokenMaker.VerifyToken(t.Token)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, PayloadKey, payload), nil
}
