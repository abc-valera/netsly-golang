package application

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/mock/mockQuery"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/mock/mockTransactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/service/passwordMaker/mockPasswordMaker"
	"github.com/abc-valera/netsly-api-golang/internal/service/taskQueuer/mockTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/service/tokenMaker/mockTokenMaker"
	"github.com/stretchr/testify/suite"
)

type SignUseCaseSuite struct {
	suite.Suite

	tx            *mockTransactioneer.Transactioneer
	userQuery     *mockQuery.User
	passwordMaker *mockPasswordMaker.PasswordMaker
	tokenMaker    *mockTokenMaker.TokenMaker
	taskQueue     *mockTaskQueuer.TaskQueuer
	userDomain    entity.User

	sign SignUseCase
}

func (suite *SignUseCaseSuite) SetupTest() {
	// Make sure to run the test in parallel
	suite.T().Parallel()

	// TODO: init everything
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
