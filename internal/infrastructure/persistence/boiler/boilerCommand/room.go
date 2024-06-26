package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type room struct {
	executor boil.ContextExecutor
}

func NewRoom(executor boil.ContextExecutor) command.IRoom {
	return &room{
		executor: executor,
	}
}

func (r room) Create(ctx context.Context, req model.Room) (model.Room, error) {
	room := sqlboiler.Room{
		ID:            req.ID,
		Name:          req.Name,
		Description:   req.Description,
		CreatedAt:     req.CreatedAt,
		CreatorUserID: req.CreatorUserID,
	}
	err := room.Insert(ctx, r.executor, boil.Infer())
	return boilerDto.NewDomainRoomWithErrHandle(&room, err)
}

func (r room) Update(ctx context.Context, id string, req command.RoomUpdate) (model.Room, error) {
	room, err := sqlboiler.FindRoom(ctx, r.executor, id)
	if err != nil {
		return model.Room{}, errutil.HandleErr(err)
	}
	if req.Description != nil {
		room.Description = *req.Description
	}
	_, err = room.Update(ctx, r.executor, boil.Infer())
	return boilerDto.NewDomainRoomWithErrHandle(room, err)
}

func (r room) Delete(ctx context.Context, id string) error {
	room, err := sqlboiler.FindRoom(ctx, r.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = room.Delete(ctx, r.executor)
	return errutil.HandleErr(err)
}
