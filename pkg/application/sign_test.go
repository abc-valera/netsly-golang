package application_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockEntity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockPasswordMaker"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockQuery"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockTokenMaker"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/mock/mockTransactioneer"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUseCase(t *testing.T) {
	type Mocks struct {
		userEntity    *mockEntity.User
		userQuery     *mockQuery.User
		tx            *mockTransactioneer.Transactioneer
		passwordMaker *mockPasswordMaker.PasswordMaker
		tokenMaker    *mockTokenMaker.TokenMaker
		taskQueue     *mockTaskQueuer.TaskQueuer
	}

	setupTest := func(t *testing.T) (*require.Assertions, Mocks, application.ISignUseCase) {
		mocks := Mocks{
			userEntity:    mockEntity.NewUser(t),
			userQuery:     mockQuery.NewUser(t),
			tx:            mockTransactioneer.NewTransactioneer(t),
			passwordMaker: mockPasswordMaker.NewPasswordMaker(t),
			tokenMaker:    mockTokenMaker.NewTokenMaker(t),
			taskQueue:     mockTaskQueuer.NewTaskQueuer(t),
		}
		return require.New(t), mocks, application.NewSignUseCase(
			mocks.userEntity,
			mocks.userQuery,
			mocks.tx,
			mocks.passwordMaker,
			mocks.tokenMaker,
			mocks.taskQueue,
		)
	}

	t.Run("SignUseCase", func(t *testing.T) {
		t.Run("SignUp", func(t *testing.T) {
			r, mocks, signUseCase := setupTest(t)

			ctx := context.Background()
			req := application.SignUpRequest{
				Username: "test",
				Email:    "test@gmail.com",
				Password: "test",
				Fullname: "test",
				Status:   "test",
			}

			mocks.tx.EXPECT().
				PerformTX(ctx, mock.Anything).
				RunAndReturn(
					func(ctx context.Context, fn func(context.Context, domain.Entities) error) error {
						return fn(ctx, domain.Entities{
							User: mocks.userEntity,
						})
					},
				)

			userCreateReq := entity.UserCreateRequest{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
				Fullname: req.Fullname,
				Status:   req.Status,
			}
			mocks.userEntity.EXPECT().Create(ctx, userCreateReq).Return(model.User{}, nil)

			sendEmail := service.Email{
				Subject: "Verification Email for Netsly!",
				Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", req.Username),
				To:      []string{req.Email},
			}
			mocks.taskQueue.EXPECT().SendEmailTask(ctx, service.Critical, sendEmail).Return(nil)

			err := signUseCase.SignUp(ctx, req)
			r.NoError(err)
		})

		t.Run("SignIn", func(t *testing.T) {
			r, mocks, signUseCase := setupTest(t)

			ctx := context.Background()
			req := application.SignInRequest{
				Email:    "test@gmail.com",
				Password: "test-test",
			}
			expected := application.SignInResponse{
				User: model.User{
					ID:             "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					HashedPassword: "test-test_hashed",
				},
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			}

			mocks.userQuery.EXPECT().GetByEmail(ctx, req.Email).Return(expected.User, nil)

			mocks.passwordMaker.EXPECT().CheckPassword(req.Password, expected.User.HashedPassword).Return(nil)

			mocks.tokenMaker.EXPECT().CreateAccessToken(expected.User.ID).Return(expected.AccessToken, nil)
			mocks.tokenMaker.EXPECT().CreateRefreshToken(expected.User.ID).Return(expected.RefreshToken, nil)

			actual, err := signUseCase.SignIn(ctx, req)
			r.NoError(err)
			r.Equal(expected, actual)
		})

		t.Run("SignRefresh", func(t *testing.T) {
			t.Run("Ok", func(t *testing.T) {
				r, mocks, signUseCase := setupTest(t)

				ctx := context.Background()
				req := "refresh_token"
				expected := "access_token"

				payloadUserID := "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
				mocks.tokenMaker.EXPECT().VerifyToken(req).Return(service.AuthPayload{
					UserID:    payloadUserID,
					IsRefresh: true,
				}, nil)

				mocks.tokenMaker.EXPECT().CreateAccessToken(payloadUserID).Return(expected, nil)

				actual, err := signUseCase.SignRefresh(ctx, req)
				r.NoError(err)
				r.Equal(expected, actual)
			})

			t.Run("ErrProvidedAccessToken", func(t *testing.T) {
				r, mocks, signUseCase := setupTest(t)

				ctx := context.Background()
				req := "refresh_token"
				expectedErr := application.ErrProvidedAccessToken

				payloadUserID := "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
				mocks.tokenMaker.EXPECT().VerifyToken(req).Return(service.AuthPayload{
					UserID:    payloadUserID,
					IsRefresh: false,
				}, nil)

				actual, actualErr := signUseCase.SignRefresh(ctx, req)
				r.Error(actualErr)
				r.Empty(actual)
				r.Equal(expectedErr, actualErr)
			})
		})
	})
}
