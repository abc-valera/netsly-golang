package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

var ErrProvidedAccessToken = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Access token provided")

type ISignUseCase interface {
	SignUp(ctx context.Context, req SignUpRequest) (model.User, error)
	SignIn(ctx context.Context, req SignInRequest) (model.User, error)
}

type signUseCase struct {
	user          entity.IUser
	transactor    entityTransactor.ITransactor
	passwordMaker service.IPasswordMaker
	taskQueue     service.ITaskQueuer
}

func NewSignUseCase(
	userEntity entity.IUser,
	transactor entityTransactor.ITransactor,
	passwordMaker service.IPasswordMaker,
	taskQueue service.ITaskQueuer,
) ISignUseCase {
	return signUseCase{
		user:          userEntity,
		transactor:    transactor,
		passwordMaker: passwordMaker,
		taskQueue:     taskQueue,
	}
}

type SignUpRequest struct {
	Username string
	Email    string
	Password string
	Fullname string
	Status   string
}

// SignUp performs user sign-up:
// it creates new user entity with unique username and email,
// creates hash of the password provided by user,
// then it sends welcome email to the users's email address,
func (uc signUseCase) SignUp(ctx context.Context, req SignUpRequest) (model.User, error) {
	var user model.User
	txFunc := func(ctx context.Context, txEntities domain.Entities) error {
		createdUser, err := txEntities.User.Create(ctx, entity.UserCreateRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Fullname: req.Fullname,
			Status:   req.Status,
		})
		if err != nil {
			return err
		}
		user = createdUser

		welcomeEmail := service.Email{
			Subject: "Verification Email for Netsly!",
			Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", req.Username),
			To:      []string{req.Email},
		}

		return uc.taskQueue.SendEmailTask(ctx, service.Critical, welcomeEmail)
	}

	if err := uc.transactor.PerformTX(ctx, txFunc); err != nil {
		return model.User{}, err
	}

	return user, nil
}

type SignInRequest struct {
	Email    string
	Password string
}

// SignIn performs user sign-in: it checks if user with provided email exists,
// then creates hash of the provided password and compares it to the hash stored in database.
// The SignIn returns user, accessToken and refreshToken.
func (s signUseCase) SignIn(ctx context.Context, req SignInRequest) (model.User, error) {
	user, err := s.user.GetByEmail(ctx, req.Email)
	if err != nil {
		return model.User{}, err
	}

	if err := s.passwordMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return model.User{}, err
	}

	return user, nil
}
