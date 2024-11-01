package application

import (
	"context"
	"fmt"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"go.opentelemetry.io/otel/trace"
)

var ErrProvidedAccessToken = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Access token provided")

type ISignUsecase interface {
	SignUp(ctx context.Context, req SignUpRequest) (model.User, error)
	SignIn(ctx context.Context, req SignInRequest) (model.User, error)
}

type signUsecase struct {
	IDependency
}

func newSignUsecase(dep IDependency) ISignUsecase {
	return signUsecase{dep}
}

type SignUpRequest struct {
	Username string
	Email    string
	Password string
	Fullname string
	Status   string
}

var welcomeEmailTemplateFunc = func(username, email string) entity.EmailSendRequest {
	return entity.EmailSendRequest{
		Subject: "Verification Email for Netsly!",
		Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", username),
		To:      []string{email},
	}
}

// SignUp performs user sign-up:
// it creates new user entity with unique username and email,
// creates hash of the password provided by user,
// then it sends welcome email to the users's email address,
func (u signUsecase) SignUp(ctx context.Context, req SignUpRequest) (model.User, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	var user model.User
	txFunc := func(ctx context.Context, txE entity.Entities, _ Usecases) error {
		createdUser, err := txE.User.Create(ctx, entity.UserCreateRequest{
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

		return u.E().Emailer.SendEmail(welcomeEmailTemplateFunc(req.Username, req.Email))
	}

	if err := u.RunInTX(ctx, txFunc); err != nil {
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
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	user, err := u.E().User.Get(ctx, model.User{Email: req.Email})
	if err != nil {
		return model.User{}, err
	}

	span.AddEvent("Check Password Start")

	if err := u.E().Passworder.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return model.User{}, err
	}

	span.AddEvent("Check Password End")

	return user, nil
}
