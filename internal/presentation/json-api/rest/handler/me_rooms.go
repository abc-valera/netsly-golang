package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/rest/dto"
)

type MeRooms struct {
	roomQuery  query.IRoom
	roomEntity entity.Room

	roomMemberEntity entity.RoomMember
}

func NewMeRooms(
	roomQuery query.IRoom,
	roomEntity entity.Room,
	roomMemberEntity entity.RoomMember,
) MeRooms {
	return MeRooms{
		roomQuery:        roomQuery,
		roomEntity:       roomEntity,
		roomMemberEntity: roomMemberEntity,
	}
}

func (h MeRooms) MeRoomsGet(ctx context.Context, ogenParams ogen.MeRoomsGetParams) (*ogen.Rooms, error) {
	domainRooms, err := h.roomQuery.GetAllByUserID(
		ctx,
		payloadUserID(ctx),
		dto.NewDomainSelectParams(&ogenParams.SelectParams),
	)
	return dto.NewRoomsResponse(domainRooms), err
}

func (h MeRooms) MeRoomsPost(ctx context.Context, req *ogen.MeRoomsPostReq) (*ogen.Room, error) {
	userID := payloadUserID(ctx)

	domainRoom, err := h.roomEntity.Create(ctx, entity.RoomCreateRequest{
		Name:        req.Name,
		CreatorID:   userID,
		Description: *dto.NewPointerString(req.Description),
	})
	return dto.NewRoomResponse(domainRoom), err
}

func (h MeRooms) MeRoomsPut(ctx context.Context, req *ogen.MeRoomsPutReq) (*ogen.Room, error) {
	domainRoom, err := h.roomEntity.Update(ctx, req.ID, entity.RoomUpdateRequest{
		Name:        dto.NewPointerString(req.Name),
		Description: dto.NewPointerString(req.Description),
	})
	return dto.NewRoomResponse(domainRoom), err
}

func (h MeRooms) MeRoomsDelete(ctx context.Context, req *ogen.MeRoomsDeleteReq) error {
	return h.roomEntity.Delete(ctx, req.ID)
}

func (h MeRooms) MeChatRoomsJoinPost(ctx context.Context, req *ogen.MeChatRoomsJoinPostReq) error {
	userID := payloadUserID(ctx)

	_, err := h.roomMemberEntity.Create(ctx, entity.RoomMemberCreateRequest{
		UserID: userID,
		RoomID: req.ID,
	})
	return err
}
