package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

var (
	ErrChatMemberNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "Chat member not found")
)

type IChatMemberQuery interface {
	GetMany(ctx context.Context, params ChatMemberSelectParams) (model.ChatMembers, error)
	GetOne(ctx context.Context, fields ChatMemberGetFields) (*model.ChatMember, error)
}

type ChatMemberSearchByFields struct {
	ChatRoomID string
	UserID     string
}

type ChatMemberOrderByFields struct {
	CreatedAt bool
}

type ChatMemberSelectParams struct {
	SearchBy ChatMemberSearchByFields
	OrderBy  ChatMemberOrderByFields
	spec.SelectParams
}

func NewChatMemberSelectParams(
	searchBy ChatMemberSearchByFields,
	orderBy ChatMemberOrderByFields,
	order string,
	limit int,
	offset int,
) (ChatMemberSelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return ChatMemberSelectParams{}, err
	}
	return ChatMemberSelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type ChatMemberGetFields struct {
	ChatRoomID string
	UserID     string
}
