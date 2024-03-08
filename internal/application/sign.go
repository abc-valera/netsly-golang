package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

var (
	ErrProvidedAccessToken = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Access token provided")
)

type SignUseCase struct {
	tx            transactioneer.ITransactioneer
	userQuery     query.IUser
	userDomain    entity.User
	passwordMaker service.IPasswordMaker
	tokenMaker    service.ITokenMaker
	messageBroker service.IMessageBroker
}

func NewSignUseCase(
	userQuery query.IUser,
	tx transactioneer.ITransactioneer,
	userDomain entity.User,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	messageBroker service.IMessageBroker,
) SignUseCase {
	return SignUseCase{
		userQuery:     userQuery,
		tx:            tx,
		userDomain:    userDomain,
		passwordMaker: passwordMaker,
		tokenMaker:    tokenMaker,
		messageBroker: messageBroker,
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
func (uc SignUseCase) SignUp(ctx context.Context, req SignUpRequest) error {
	txFunc := func(ctx context.Context, commands domain.Commands) error {
		if _, err := uc.userDomain.Create(ctx, entity.UserCreateRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Fullname: req.Fullname,
			Status:   req.Status,
		}); err != nil {
			return err
		}

		welcomeEmail := service.Email{
			Subject: "Verification Email for Netsly!",
			Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", req.Username),
			To:      []string{req.Email},
		}

		return uc.messageBroker.SendEmailTask(ctx, service.Critical, welcomeEmail)
	}

	return uc.tx.PerformTX(ctx, txFunc)
}

type SignInRequest struct {
	Email    string
	Password string
}

type SignInResponse struct {
	User         model.User
	AccessToken  string
	RefreshToken string
}

// SignIn performs user sign-in: it checks if user with provided email exists,
// then creates hash of the provided password and compares it to the hash stored in database.
// The SignIn returns user, accessToken and refreshToken.
func (s SignUseCase) SignIn(ctx context.Context, req SignInRequest) (SignInResponse, error) {
	user, err := s.userQuery.GetByEmail(ctx, req.Email)
	if err != nil {
		return SignInResponse{}, err
	}

	if err := s.passwordMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return SignInResponse{}, err
	}

	access, _, err := s.tokenMaker.CreateAccessToken(user.ID)
	if err != nil {
		return SignInResponse{}, err
	}
	refresh, _, err := s.tokenMaker.CreateRefreshToken(user.ID)
	if err != nil {
		return SignInResponse{}, err
	}

	return SignInResponse{
		User:         user,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// SignRefresh exchages given refresh token for the access token for the same user.
func (s SignUseCase) SignRefresh(c context.Context, refreshToken string) (string, error) {
	payload, err := s.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	if !payload.IsRefresh {
		return "", ErrProvidedAccessToken
	}

	access, _, err := s.tokenMaker.CreateAccessToken(payload.UserID)
	if err != nil {
		return "", err
	}

	return access, nil
}
