package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

var (
	ErrChatRoomNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "ChatRoom not found")
)

type IChatRoomQuery interface {
	GetAll(ctx context.Context, params ChatRoomSelectParams) (model.ChatRooms, error)
	GetOne(ctx context.Context, fields ChatRoomGetFields) (*model.ChatRoom, error)
}

type ChatRoomSearchByFields struct {
	UserID string
	Title  string
}

type ChatRoomOrderByFields struct {
	Title     bool
	CreatedAt bool
}

type ChatRoomSelectParams struct {
	SearchBy ChatRoomSearchByFields
	OrderBy  ChatRoomOrderByFields
	spec.SelectParams
}

func NewChatRoomSelectParams(
	searchBy ChatRoomSearchByFields,
	orderBy ChatRoomOrderByFields,
	order string,
	limit int,
	offset int,
) (ChatRoomSelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return ChatRoomSelectParams{}, err
	}
	return ChatRoomSelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type ChatRoomGetFields struct {
	ID     string
	UserID string
	Title  string
}
