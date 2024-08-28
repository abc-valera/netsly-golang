package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type room struct {
	db bun.IDB
}

func NewRoom(db bun.IDB) command.IRoom {
	return &room{
		db: db,
	}
}

func (c room) Create(ctx context.Context, req command.RoomCreateRequest) (model.Room, error) {
	room := bunSqliteDto.Room{
		ID:          req.Room.ID,
		Name:        req.Room.Name,
		Description: req.Room.Description,
		CreatedAt:   req.Room.CreatedAt,
		UpdatedAt:   req.Room.UpdatedAt,
		DeletedAt:   req.Room.DeletedAt,

		CreatorUserID: req.CreatorUserID,
	}

	res, err := c.db.NewInsert().Model(&room).Exec(ctx)
	return room.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c room) Update(ctx context.Context, id string, req command.RoomUpdateRequest) (model.Room, error) {
	room := bunSqliteDto.Room{
		ID: id,
	}
	var columns []string

	if req.Description != nil {
		room.Description = *req.Description
		columns = append(columns, "description")
	}

	if len(columns) == 0 {
		return model.Room{}, nil
	}

	res, err := c.db.NewUpdate().Model(&room).Column(columns...).WherePK().Exec(ctx)
	return room.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c room) Delete(ctx context.Context, id string) error {
	room := bunSqliteDto.Room{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&room).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
