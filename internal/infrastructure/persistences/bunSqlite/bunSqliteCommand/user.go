package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type user struct {
	db bun.IDB
}

func NewUser(db bun.IDB) command.IUser {
	return &user{
		db: db,
	}
}

func (c user) Create(ctx context.Context, req model.User) (model.User, error) {
	user := bunSqliteDto.User{
		ID:             req.ID,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		UpdatedAt:      req.UpdatedAt,
		CreatedAt:      req.CreatedAt,
		DeletedAt:      req.DeletedAt,
	}
	res, err := c.db.NewInsert().Model(&user).Exec(ctx)
	return user.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c user) Update(ctx context.Context, id string, req command.UserUpdateRequest) (model.User, error) {
	user := bunSqliteDto.User{
		ID: id,
	}
	var columns []string

	if req.HashedPassword != nil {
		user.HashedPassword = *req.HashedPassword
		columns = append(columns, "hashed_password")
	}
	if req.Fullname != nil {
		user.Fullname = *req.Fullname
		columns = append(columns, "fullname")
	}
	if req.Status != nil {
		user.Status = *req.Status
		columns = append(columns, "status")
	}

	if len(columns) == 0 {
		return model.User{}, nil
	}

	res, err := c.db.NewUpdate().Model(&user).Column(columns...).WherePK().Exec(ctx)
	return user.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c user) Delete(ctx context.Context, id string) error {
	user := bunSqliteDto.User{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&user).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
