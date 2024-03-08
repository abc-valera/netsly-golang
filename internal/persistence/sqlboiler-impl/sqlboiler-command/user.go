package sqlboilercommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type user struct {
	executor boil.ContextExecutor
}

func newUser(executor boil.ContextExecutor) command.IUser {
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
		Fullname:       null.NewString(req.Fullname, req.Fullname != ""),
		Status:         null.NewString(req.Status, req.Status != ""),
		CreatedAt:      req.CreatedAt,
	}
	err := user.Insert(ctx, u.executor, boil.Infer())
	return dto.ToDomainUserWithErrHandle(&user, err)
}

func (u user) Update(ctx context.Context, id string, req command.UserUpdate) (model.User, error) {
	user, err := sqlboiler.FindUser(ctx, u.executor, id)
	if err != nil {
		return model.User{}, errors.HandleErr(err)
	}
	if req.HashedPassword != nil {
		user.HashedPassword = *req.HashedPassword
	}
	if req.Fullname != nil {
		user.Fullname = null.NewString(*req.Fullname, *req.Fullname != "")
	}
	if req.Status != nil {
		user.Status = null.NewString(*req.Status, *req.Status != "")
	}
	_, err = user.Update(ctx, u.executor, boil.Infer())
	return dto.ToDomainUserWithErrHandle(user, err)
}

func (u user) Delete(ctx context.Context, id string) error {
	user, err := sqlboiler.FindUser(ctx, u.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = user.Delete(ctx, u.executor)
	return errors.HandleErr(err)
}
