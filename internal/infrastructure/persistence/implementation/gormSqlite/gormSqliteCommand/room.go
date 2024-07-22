package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type room struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) command.IRoom {
	return &room{
		db: db,
	}
}

func (c room) Create(ctx context.Context, req command.RoomCreateRequest) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room := gormSqliteDto.Room{
		ID:            req.Room.ID,
		Name:          req.Room.Name,
		Description:   req.Room.Description,
		CreatedAt:     req.Room.CreatedAt,
		UpdatedAt:     req.Room.UpdatedAt,
		DeletedAt:     req.Room.DeletedAt,
		CreatorUserID: req.CreatorUserID,
	}
	res := c.db.Create(&room)
	return gormSqliteDto.NewDomainRoom(room), gormSqliteErrutil.HandleCommandResult(res)
}

func (c room) Update(ctx context.Context, id string, req command.RoomUpdateRequest) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var room gormSqliteDto.Room
	queryRes := c.db.Where("id = ?", id).First(&room)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.Room{}, err
	}

	if req.Description != nil {
		room.Description = *req.Description
	}

	updateRes := c.db.Save(&room)
	return gormSqliteDto.NewDomainRoom(room), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c room) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room := gormSqliteDto.Room{
		ID: id,
	}
	res := c.db.Delete(&room)
	return gormSqliteErrutil.HandleCommandResult(res)
}
