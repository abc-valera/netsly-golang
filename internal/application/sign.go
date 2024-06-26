package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/domain/transactor"
)

var ErrProvidedAccessToken = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Access token provided")

type ISignUsecase interface {
	SignUp(ctx context.Context, req SignUpRequest) (model.User, error)
	SignIn(ctx context.Context, req SignInRequest) (model.User, error)
}

type signUsecase struct {
	user          entity.IUser
	transactor    transactor.ITransactor
	passwordMaker service.IPasswordMaker
	taskQueue     service.ITaskQueuer
}

func NewSignUsecase(
	userEntity entity.IUser,
	transactor transactor.ITransactor,
	passwordMaker service.IPasswordMaker,
	taskQueue service.ITaskQueuer,
) ISignUsecase {
	return signUsecase{
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
func (u signUsecase) SignUp(ctx context.Context, req SignUpRequest) (model.User, error) {
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

		return u.taskQueue.SendEmailTask(ctx, service.Critical, welcomeEmail)
	}

	if err := u.transactor.PerformTX(ctx, txFunc); err != nil {
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
func (u signUsecase) SignIn(ctx context.Context, req SignInRequest) (model.User, error) {
	user, err := u.user.GetByEmail(ctx, req.Email)
	if err != nil {
		return model.User{}, err
	}

	if err := u.passwordMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return model.User{}, err
	}

	return user, nil
}
