package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

var (
	ErrProvidedAccessToken = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Access token provided")
)

type SignUseCase struct {
	tx            transactioneer.ITransactioneer
	userQuery     query.IUserQuery
	userDomain    domain.UserDomain
	passwordMaker service.IPasswordMaker
	tokenMaker    service.ITokenMaker
	messageBroker service.IMessageBroker
}

func NewSignUseCase(
	userQuery query.IUserQuery,
	tx transactioneer.ITransactioneer,
	userDomain domain.UserDomain,
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
	txFunc := func(ctx context.Context) error {
		if err := uc.userDomain.Create(ctx, domain.UserCreateRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Fullname: req.Fullname,
			Status:   req.Status,
		}); err != nil {
			return err
		}

		welcomeEmail := service.Email{
			Subject: "Verification Email for Flugo!",
			Content: fmt.Sprintf(`%s, congrats with joining the Flugo community!<br/>`, user.Username),
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
	User         *model.User
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
