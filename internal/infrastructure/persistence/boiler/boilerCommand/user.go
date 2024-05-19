package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errors"
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
		return model.User{}, errors.HandleErr(err)
	}
	if req.HashedPassword.IsPresent() {
		user.HashedPassword = req.HashedPassword.Value()
	}
	if req.Fullname.IsPresent() {
		user.Fullname = req.Fullname.Value()
	}
	if req.Status.IsPresent() {
		user.Status = req.Status.Value()
	}
	_, err = user.Update(ctx, u.executor, boil.Infer())
	return boilerDto.NewDomainUserWithErrHandle(user, err)
}

func (u user) Delete(ctx context.Context, id string) error {
	user, err := sqlboiler.FindUser(ctx, u.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = user.Delete(ctx, u.executor)
	return errors.HandleErr(err)
}
