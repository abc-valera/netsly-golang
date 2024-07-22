package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
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

func (c room) Create(ctx context.Context, req command.RoomCreateRequest) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room := sqlboiler.Room{
		ID:          req.Room.ID,
		Name:        req.Room.Name,
		Description: req.Room.Description,
		CreatedAt:   req.Room.CreatedAt,
		UpdatedAt:   req.Room.UpdatedAt,
		DeletedAt:   req.Room.DeletedAt,

		CreatorUserID: req.CreatorUserID,
	}
	err := room.Insert(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainRoom(&room), boilerSqliteErrutil.HandleErr(err)
}

func (c room) Update(ctx context.Context, id string, req command.RoomUpdateRequest) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room, err := sqlboiler.FindRoom(ctx, c.executor, id)
	if err != nil {
		return model.Room{}, boilerSqliteErrutil.HandleErr(err)
	}
	if req.Description != nil {
		room.Description = *req.Description
	}
	_, err = room.Update(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainRoom(room), boilerSqliteErrutil.HandleErr(err)
}

func (c room) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room, err := sqlboiler.FindRoom(ctx, c.executor, id)
	if err != nil {
		return boilerSqliteErrutil.HandleErr(err)
	}
	_, err = room.Delete(ctx, c.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
