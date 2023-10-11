package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

var (
	ErrProvidedAccessToken = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Access token provided")
)

type SignUseCase struct {
	userRepo      repository.IUserRepository
	passwordMaker service.IPasswordMaker
	tokenMaker    service.ITokenMaker
	msgBroker     service.IMessageBroker
}

func NewSignUseCase(
	userRepo repository.IUserRepository,
	passwordMaker service.IPasswordMaker,
	tokenMaker service.ITokenMaker,
	msgBroker service.IMessageBroker,
) SignUseCase {
	return SignUseCase{
		userRepo:      userRepo,
		passwordMaker: passwordMaker,
		tokenMaker:    tokenMaker,
		msgBroker:     msgBroker,
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
func (s SignUseCase) SignUp(ctx context.Context, req SignUpRequest) error {
	hashedPassword, err := s.passwordMaker.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user, err := entity.NewUser(req.Username, req.Email, hashedPassword, req.Fullname, req.Status)
	if err != nil {
		return err
	}

	txFunc := func(ctx context.Context) error {
		if err := s.userRepo.Create(ctx, user); err != nil {
			return err
		}

		welcomeEmail := service.Email{
			Subject: "Verification Email for Flugo!",
			Content: fmt.Sprintf(`%s, congrats with joining the Flugo community!<br/>`, user.Username),
			To:      []string{user.Email},
		}

		return s.msgBroker.SendEmailTask(ctx, service.Critical, welcomeEmail)
	}

	return s.userRepo.PerformTX(ctx, txFunc)
}

type SignInRequest struct {
	Email    string
	Password string
}

// SignIn performs user sign-in: it checks if user with provided email exists,
// then creates hash of the provided password and compares it to the hash stored in database.
// The SignIn returns user, accessToken and refreshToken.
func (s SignUseCase) SignIn(ctx context.Context, req SignInRequest) (*entity.User, string, string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", "", err
	}

	if err := s.passwordMaker.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return nil, "", "", err
	}

	access, _, err := s.tokenMaker.CreateAccessToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}
	refresh, _, err := s.tokenMaker.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}

	return user, access, refresh, nil
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
