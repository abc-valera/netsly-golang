package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
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
	domainRooms, err := h.room.GetAllByUserID(
		ctx,
		payloadUserID(ctx),
		restDto.NewDomainSelector(&ogenParams.Selector),
	)
	return restDto.NewRoomsResponse(domainRooms), err
}

func (h MeRooms) MeRoomsPost(ctx context.Context, req *ogen.MeRoomsPostReq) (*ogen.Room, error) {
	userID := payloadUserID(ctx)

	domainRoom, err := h.room.Create(ctx, entity.RoomCreateRequest{
		Name:          req.Name,
		CreatorUserID: userID,
		Description:   *restDto.NewPointerString(req.Description),
	})
	return restDto.NewRoomResponse(domainRoom), err
}

func (h MeRooms) MeRoomsPut(ctx context.Context, req *ogen.MeRoomsPutReq) (*ogen.Room, error) {
	domainRoom, err := h.room.Update(ctx, req.ID, entity.RoomUpdateRequest{
		Name:        restDto.NewPointerString(req.Name),
		Description: restDto.NewPointerString(req.Description),
	})
	return restDto.NewRoomResponse(domainRoom), err
}

func (h MeRooms) MeRoomsDelete(ctx context.Context, req *ogen.MeRoomsDeleteReq) error {
	return h.room.Delete(ctx, req.ID)
}

func (h MeRooms) MeChatRoomsJoinPost(ctx context.Context, req *ogen.MeChatRoomsJoinPostReq) error {
	userID := payloadUserID(ctx)

	_, err := h.roomMember.Create(ctx, entity.RoomMemberCreateRequest{
		UserID: userID,
		RoomID: req.ID,
	})
	return err
}
