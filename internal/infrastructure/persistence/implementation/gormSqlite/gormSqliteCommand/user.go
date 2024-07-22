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

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) command.IUser {
	return &user{
		db: db,
	}
}

func (c user) Create(ctx context.Context, req model.User) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	user := gormSqliteDto.User{
		ID:             req.ID,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
		DeletedAt:      req.DeletedAt,
	}
	res := c.db.Create(&user)
	return gormSqliteDto.NewDomainUser(user), gormSqliteErrutil.HandleCommandResult(res)
}

func (c user) Update(ctx context.Context, id string, req command.UserUpdateRequest) (model.User, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var user gormSqliteDto.User
	queryRes := c.db.Where("id = ?", id).First(&user)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.User{}, err
	}

	if req.HashedPassword != nil {
		user.HashedPassword = *req.HashedPassword
	}
	if req.Fullname != nil {
		user.Fullname = *req.Fullname
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	updateRes := c.db.Save(&user)
	return gormSqliteDto.NewDomainUser(user), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c user) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	user := gormSqliteDto.User{
		ID: id,
	}
	res := c.db.Delete(&user)
	return gormSqliteErrutil.HandleCommandResult(res)
}
