package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

var (
	ErrChatMessageNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "ChatMessage not found")
)

type IChatMessageQuery interface {
	GetMany(ctx context.Context, params ChatMessageSelectParams) (model.ChatMessages, error)
	GetOne(ctx context.Context, fields ChatRoomGetFields) (*model.ChatMessage, error)
}

type ChatMessageSearchFields struct {
	UserID string
	Text   string
}

type ChatMessageOrderByFields struct {
	CreatedAt bool
}

type ChatMessageSelectParams struct {
	SearchBy ChatMessageSearchFields
	OrderBy  ChatMessageOrderByFields
	spec.SelectParams
}

func NewChatMessageSelectParams(
	searchBy ChatMessageSearchFields,
	orderBy ChatMessageOrderByFields,
	order string,
	limit int,
	offset int,
) (ChatMessageSelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return ChatMessageSelectParams{}, err
	}
	return ChatMessageSelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type ChatMessageGetFields struct {
	ID     string
	UserID string
	Text   string
}
