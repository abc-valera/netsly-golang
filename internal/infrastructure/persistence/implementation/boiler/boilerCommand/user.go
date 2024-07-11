package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type user struct {
	executor boil.ContextExecutor
}

func NewUser(executor boil.ContextExecutor) command.IUser {
	return &user{
		executor: executor,
	}
}

func (u user) Create(ctx context.Context, req model.User) (model.User, error) {
	user := sqlboiler.User{
		ID:             req.ID,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		CreatedAt:      req.CreatedAt,
	}
	err := user.Insert(ctx, u.executor, boil.Infer())
	return boilerDto.NewDomainUserWithErrHandle(&user, err)
}

func (u user) Update(ctx context.Context, id string, req command.UserUpdate) (model.User, error) {
	user, err := sqlboiler.FindUser(ctx, u.executor, id)
	if err != nil {
		return model.User{}, errutil.HandleErr(err)
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
	_, err = user.Update(ctx, u.executor, boil.Infer())
	return boilerDto.NewDomainUserWithErrHandle(user, err)
}

func (u user) Delete(ctx context.Context, id string) error {
	user, err := sqlboiler.FindUser(ctx, u.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = user.Delete(ctx, u.executor)
	return errutil.HandleErr(err)
}
