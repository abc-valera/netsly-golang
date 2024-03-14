package application_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/mock/mockCommand"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/mock/mockQuery"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/mock/mockTransactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/service/passwordMaker/mockPasswordMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/taskQueuer/mockTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/service/timeMaker/mockTimeMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/tokenMaker/mockTokenMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/uuidMaker/mockUuidMaker"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignUseCaseSuite struct {
	suite.Suite

	tx            *mockTransactioneer.Transactioneer
	userQuery     *mockQuery.User
	passwordMaker *mockPasswordMaker.PasswordMaker
	tokenMaker    *mockTokenMaker.TokenMaker
	taskQueue     *mockTaskQueuer.TaskQueuer

	userCommand *mockCommand.User
	uuidMaker   *mockUuidMaker.UuidMaker
	timeMaker   *mockTimeMaker.TimeMaker
	userEntity  entity.User

	sign application.SignUseCase
}

func (suite *SignUseCaseSuite) SetupTest() {
	// Make sure to run the test in parallel
	suite.T().Parallel()

	suite.tx = mockTransactioneer.NewTransactioneer(suite.T())
	suite.userQuery = mockQuery.NewUser(suite.T())
	suite.passwordMaker = mockPasswordMaker.NewPasswordMaker(suite.T())
	suite.tokenMaker = mockTokenMaker.NewTokenMaker(suite.T())
	suite.taskQueue = mockTaskQueuer.NewTaskQueuer(suite.T())

	suite.userCommand = mockCommand.NewUser(suite.T())
	suite.uuidMaker = mockUuidMaker.NewUuidMaker(suite.T())
	suite.timeMaker = mockTimeMaker.NewTimeMaker(suite.T())
	suite.userEntity = entity.NewUser(
		suite.userCommand,
		suite.userQuery,
		suite.uuidMaker,
		suite.timeMaker,
		suite.passwordMaker,
	)

	suite.sign = application.NewSignUseCase(
		suite.userEntity,
		suite.userQuery,
		suite.tx,
		suite.passwordMaker,
		suite.tokenMaker,
		suite.taskQueue,
	)
}

func (suite *SignUseCaseSuite) TestSignUp() {
	// Prepare the begging request
	req := application.SignUpRequest{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "test",
		Fullname: "test",
		Status:   "test",
	}

	// Expectations for transactioneer
	suite.tx.EXPECT().
		PerformTX(context.Background(), mock.Anything).
		RunAndReturn(func(ctx context.Context, f func(context.Context, domain.Entities) error) error {
			return f(ctx, domain.Entities{
				User: suite.userEntity,
			})
		})

	// Prepare the user command request
	userCommandReq := model.User{
		ID:             "a5ed4a29-8160-465e-b46c-75038c43f4f3",
		Username:       "test",
		Email:          "test@gmail.com",
		HashedPassword: "test_hashed",
		Fullname:       "test",
		Status:         "test",
		CreatedAt:      time.Now(),
	}

	// Expectations for user entity
	suite.passwordMaker.EXPECT().HashPassword(req.Password).Return(userCommandReq.HashedPassword, nil)
	suite.uuidMaker.EXPECT().NewUUID().Return(userCommandReq.ID)
	suite.timeMaker.EXPECT().Now().Return(userCommandReq.CreatedAt)
	suite.userCommand.EXPECT().Create(context.Background(), userCommandReq).Return(userCommandReq, nil)

	// Prepare the email for task queue
	email := service.Email{
		Subject: "Verification Email for Netsly!",
		Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", req.Username),
		To:      []string{req.Email},
	}

	// Expectations for task queue
	suite.taskQueue.EXPECT().SendEmailTask(context.Background(), service.Critical, email).Return(nil)

	// Run the test
	err := suite.sign.SignUp(context.Background(), req)
	suite.NoError(err)
}

func (suite *SignUseCaseSuite) TestSignIn() {
	// Prepare the begging request
	req := application.SignInRequest{
		Email:    "test@gmail.com",
		Password: "test",
	}

	userQueryReturn := model.User{
		ID:             "a5ed4a29-8160-465e-b46c-75038c43f4f3",
		HashedPassword: "test_hashed",
	}
	suite.userQuery.EXPECT().GetByEmail(context.Background(), req.Email).Return(userQueryReturn, nil)

	suite.passwordMaker.EXPECT().CheckPassword(req.Password, userQueryReturn.HashedPassword).Return(nil)

	accessToken := "access_token"
	suite.tokenMaker.EXPECT().CreateAccessToken(userQueryReturn.ID).Return(accessToken, service.AuthPayload{}, nil)
	refreshToken := "refresh_token"
	suite.tokenMaker.EXPECT().CreateRefreshToken(userQueryReturn.ID).Return(refreshToken, service.AuthPayload{}, nil)

	res, err := suite.sign.SignIn(context.Background(), req)
	suite.NoError(err)
	suite.Equal(userQueryReturn, res.User)
	suite.Equal(accessToken, res.AccessToken)
	suite.Equal(refreshToken, res.RefreshToken)
}

func (suite *SignUseCaseSuite) TestSignRefresh() {
	reqRefreshToken := "refresh_token"

	payload := service.AuthPayload{
		UserID:    "a5ed4a29-8160-465e-b46c-75038c43f4f3",
		IsRefresh: true,
	}
	suite.tokenMaker.EXPECT().VerifyToken(reqRefreshToken).Return(payload, nil)

	accessToken := "access_token"
	suite.tokenMaker.EXPECT().CreateAccessToken(payload.UserID).Return(accessToken, service.AuthPayload{}, nil)

	res, err := suite.sign.SignRefresh(context.Background(), reqRefreshToken)
	suite.NoError(err)
	suite.Equal(accessToken, res)
}

func TestSignUseCaseSuite(t *testing.T) {
	suite.Run(t, new(SignUseCaseSuite))
}
