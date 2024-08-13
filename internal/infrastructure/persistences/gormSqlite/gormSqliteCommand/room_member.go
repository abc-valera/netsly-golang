package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type roomMember struct {
	db *gorm.DB
}

func NewRoomMember(db *gorm.DB) command.IRoomMember {
	return &roomMember{
		db: db,
	}
}

func (c roomMember) Create(ctx context.Context, req command.RoomMemberCreateRequest) (model.RoomMember, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMember := gormSqliteDto.RoomMember{
		CreatedAt: req.CreatedAt,
		DeletedAt: req.DeletedAt,
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	}
	res := c.db.Create(&roomMember)
	return gormSqliteDto.NewDomainRoomMember(roomMember), gormSqliteErrutil.HandleCommandResult(res)
}

func (c roomMember) Delete(ctx context.Context, userID string, roomID string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMember := gormSqliteDto.RoomMember{
		UserID: userID,
		RoomID: roomID,
	}
	res := c.db.Delete(&roomMember)
	return gormSqliteErrutil.HandleCommandResult(res)
}
