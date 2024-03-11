package application

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/stretchr/testify/suite"
)

type SignUseCaseSuite struct {
	suite.Suite

	tx            *transactioneer.MockTransactioneer
	userQuery     *query.MockUser
	passwordMaker *service.MockPasswordMaker
	tokenMaker    *service.MockTokenMaker
	taskQueue     *service.MockTaskQueuer
	userDomain    entity.User

	sign SignUseCase
}

func (suite *SignUseCaseSuite) SetupTest() {
	// Make sure to run the test in parallel
	suite.T().Parallel()

	suite.tx = transactioneer.NewMockTransactioneer(suite.T())
	suite.userQuery = query.NewMockUser(suite.T())
	suite.passwordMaker = service.NewMockPasswordMaker(suite.T())
	suite.tokenMaker = service.NewMockTokenMaker(suite.T())
	suite.taskQueue = service.NewMockTaskQueuer(suite.T())
	suite.userDomain = entity.User{}

	suite.sign = NewSignUseCase(
		suite.userQuery,
		suite.tx,
		suite.userDomain,
		suite.passwordMaker,
		suite.tokenMaker,
		suite.taskQueue,
	)
}

func (suite *SignUseCaseSuite) TestSignUp() {
}

func (suite *SignUseCaseSuite) TestSignIn() {
}

func (suite *SignUseCaseSuite) TestSignRefresh() {
}

func TestSignUseCaseSuite(t *testing.T) {
	suite.Run(t, new(SignUseCaseSuite))
}
