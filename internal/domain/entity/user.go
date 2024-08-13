package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
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
	command command.IUser
	query.IUser
	passMaker service.IPassworder
}

func NewUser(
	command command.IUser,
	query query.IUser,
	passMaker service.IPassworder,
) IUser {
	return user{
		command:   command,
		IUser:     query,
		passMaker: passMaker,
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

	hashedPassword, err := e.passMaker.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	span.AddEvent("Hashing Password End")

	return e.command.Create(ctx, model.User{
		ID:             uuid.New().String(),
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Fullname:       req.Fullname,
		Status:         req.Status,
		CreatedAt:      time.Now(),
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
		HashedPassword: nil,
		Fullname:       req.Fullname,
		Status:         req.Status,
	}

	if req.Password != nil {
		span.AddEvent("Hashing Password Start")

		hashedPassword, err := e.passMaker.HashPassword(*req.Password)
		if err != nil {
			return model.User{}, err
		}
		updateReq.HashedPassword = &hashedPassword

		span.AddEvent("Hashing Password End")
	}

	return e.command.Update(ctx, userID, updateReq)
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

	if err := e.passMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	span.AddEvent("Checking Password End")

	return e.command.Delete(ctx, userID)
}
