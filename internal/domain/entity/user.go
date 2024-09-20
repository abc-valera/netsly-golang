package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IUser interface {
	Create(ctx context.Context, req UserCreateRequest) (model.User, error)
	Update(ctx context.Context, userID string, req UserUpdateRequest) (model.User, error)
	Delete(ctx context.Context, userID string, req UserDeleteRequest) error

	query.IUser
}

type user struct {
	IDependency

	query.IUser
}

func newUser(dep IDependency) IUser {
	return user{
		IDependency: dep,

		IUser: dep.Q().User,
	}
}

type UserCreateRequest struct {
	Username string `validate:"required,min=2,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=2,max=32"`
	Fullname string `validate:"max=64"`
	Status   string `validate:"max=128"`
}

func (e user) Create(ctx context.Context, req UserCreateRequest) (model.User, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.User{}, err
	}

	span.AddEvent("Hashing Password Start")

	hashedPassword, err := e.E().Passworder.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	span.AddEvent("Hashing Password End")

	return e.C().User.Create(ctx, model.User{
		ID:             uuid.New().String(),
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		CreatedAt:      time.Now().Truncate(time.Millisecond),
	})
}

type UserUpdateRequest struct {
	Password *string `validate:"min=2,max=32"`
	Fullname *string `validate:"max=64"`
	Status   *string `validate:"max=128"`
}

func (e user) Update(ctx context.Context, userID string, req UserUpdateRequest) (model.User, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.User{}, err
	}

	updateReq := command.UserUpdateRequest{
		UpdatedAt: time.Now().Truncate(time.Millisecond),

		HashedPassword: nil,
		Fullname:       req.Fullname,
		Status:         req.Status,
	}

	if req.Password != nil {
		span.AddEvent("Hashing Password Start")

		hashedPassword, err := e.E().Passworder.HashPassword(*req.Password)
		if err != nil {
			return model.User{}, err
		}
		updateReq.HashedPassword = &hashedPassword

		span.AddEvent("Hashing Password End")
	}

	return e.C().User.Update(
		ctx,
		model.User{ID: userID},
		updateReq,
	)
}

type UserDeleteRequest struct {
	Password string
}

func (e user) Delete(ctx context.Context, userID string, req UserDeleteRequest) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return err
	}

	user, err := e.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	span.AddEvent("Checking Password Start")

	if err := e.E().Passworder.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	span.AddEvent("Checking Password End")

	return e.C().User.Delete(ctx, model.User{ID: userID})
}
