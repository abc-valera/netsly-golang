package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
	"go.opentelemetry.io/otel/trace"
)

type MeRooms struct {
	room       entity.IRoom
	roomMember entity.IRoomMember
}

func NewMeRooms(
	room entity.IRoom,
	roomMember entity.IRoomMember,
) MeRooms {
	return MeRooms{
		room:       room,
		roomMember: roomMember,
	}
}

func (h MeRooms) MeRoomsGet(ctx context.Context, ogenParams ogen.MeRoomsGetParams) (*ogen.Rooms, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	domainRooms, err := h.room.GetAllByUserID(
		ctx,
		userID,
		restDto.NewDomainSelector(&ogenParams.Selector),
	)
	return restDto.NewRooms(domainRooms), err
}

func (h MeRooms) MeRoomsPost(ctx context.Context, req *ogen.MeRoomsPostReq) (*ogen.Room, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	domainRoom, err := h.room.Create(ctx, entity.RoomCreateRequest{
		Name:          req.Name,
		CreatorUserID: userID,
		Description:   req.Description.Value,
	})
	return restDto.NewRoom(domainRoom), err
}

func (h MeRooms) MeRoomsPut(ctx context.Context, req *ogen.MeRoomsPutReq) (*ogen.Room, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	domainRoom, err := h.room.Update(ctx, req.ID, entity.RoomUpdateRequest{
		Name:        restDto.NewDomainOptionalString(req.Name),
		Description: restDto.NewDomainOptionalString(req.Description),
	})
	return restDto.NewRoom(domainRoom), err
}

func (h MeRooms) MeRoomsDelete(ctx context.Context, req *ogen.MeRoomsDeleteReq) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return h.room.Delete(ctx, req.ID)
}

func (h MeRooms) MeChatRoomsJoinPost(ctx context.Context, req *ogen.MeChatRoomsJoinPostReq) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return err
	}

	if _, err := h.roomMember.Create(ctx, entity.RoomMemberCreateRequest{
		UserID: userID,
		RoomID: req.ID,
	}); err != nil {
		return err
	}

	return nil
}
