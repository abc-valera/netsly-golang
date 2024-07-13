package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
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

func (c room) Create(ctx context.Context, req model.Room) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room := sqlboiler.Room{
		ID:            req.ID,
		Name:          req.Name,
		Description:   req.Description,
		CreatedAt:     req.CreatedAt,
		CreatorUserID: req.CreatorUserID,
	}
	err := room.Insert(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainRoomWithErrHandle(&room, err)
}

func (c room) Update(ctx context.Context, id string, req command.RoomUpdate) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room, err := sqlboiler.FindRoom(ctx, c.executor, id)
	if err != nil {
		return model.Room{}, errutil.HandleErr(err)
	}
	if req.Description != nil {
		room.Description = *req.Description
	}
	_, err = room.Update(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainRoomWithErrHandle(room, err)
}

func (c room) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room, err := sqlboiler.FindRoom(ctx, c.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = room.Delete(ctx, c.executor)
	return errutil.HandleErr(err)
}
