package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type SignUseCase struct {
	userRepo      repository.UserRepository
	passwordMaker service.PasswordMaker
	tokenMaker    service.TokenMaker
	msgBroker     service.MessageBroker
}

func NewSignUseCase(
	userRepo repository.UserRepository,
	passwordMaker service.PasswordMaker,
	tokenMaker service.TokenMaker,
	msgBroker service.MessageBroker,
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

// SignIn performs user sign-in: it checks if user with provided email exists,
// then creates hash of the provided password and compares it to the hash stored in database.
// The SignIn returns user, accessToken and refreshToken.
func (s SignUseCase) SignIn(ctx context.Context, email, password string) (*entity.User, string, string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", "", err
	}

	if err := s.passwordMaker.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, "", "", err
	}

	access, _, err := s.tokenMaker.CreateAccessToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}
	refresh, _, err := s.tokenMaker.CreateRefreshToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}

	return user, access, refresh, nil
}

// SignRefresh exchages given refresh token for the access token for the same user.
func (s SignUseCase) SignRefresh(c context.Context, refreshToken string) (string, error) {
	// TODO: add implementation
	return "", nil
}
