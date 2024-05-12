package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
	"github.com/volatiletech/null/v8"
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
		Description:   null.NewString(req.Description, req.Description != ""),
		CreatedAt:     req.CreatedAt,
		CreatorUserID: req.CreatorUserID,
	}
	err := room.Insert(ctx, r.executor, boil.Infer())
	return dto.ToDomainRoomWithErrHandle(&room, err)
}

func (r room) Update(ctx context.Context, id string, req command.RoomUpdate) (model.Room, error) {
	room, err := sqlboiler.FindRoom(ctx, r.executor, id)
	if err != nil {
		return model.Room{}, errors.HandleErr(err)
	}
	if req.Description != nil {
		room.Description = null.NewString(*req.Description, *req.Description != "")
	}
	_, err = room.Update(ctx, r.executor, boil.Infer())
	return dto.ToDomainRoomWithErrHandle(room, err)
}

func (r room) Delete(ctx context.Context, id string) error {
	room, err := sqlboiler.FindRoom(ctx, r.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = room.Delete(ctx, r.executor)
	return errors.HandleErr(err)
}
