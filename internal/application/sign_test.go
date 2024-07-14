package application_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/abc-valera/netsly-api-golang/gen/mock/mockEntity"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockEntityTransactor"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockPasswordMaker"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockQuery"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUsecase(t *testing.T) {
	type Mocks struct {
		userEntity    *mockEntity.User
		userQuery     *mockQuery.User
		transactor    *mockEntityTransactor.Transactor
		passwordMaker *mockPasswordMaker.PasswordMaker
		taskQueue     *mockTaskQueuer.TaskQueuer
	}

	setupTest := func(t *testing.T) (*require.Assertions, Mocks, application.ISignUsecase) {
		mocks := Mocks{
			userEntity:    mockEntity.NewUser(t),
			userQuery:     mockQuery.NewUser(t),
			transactor:    mockEntityTransactor.NewTransactor(t),
			passwordMaker: mockPasswordMaker.NewPasswordMaker(t),
			taskQueue:     mockTaskQueuer.NewTaskQueuer(t),
		}
		return require.New(t), mocks, application.NewSignUsecase(
			mocks.userEntity,
			mocks.transactor,
			mocks.passwordMaker,
			mocks.taskQueue,
		)
	}

	t.Run("SignUsecase", func(t *testing.T) {
		t.Run("SignUp", func(t *testing.T) {
			r, mocks, signUsecase := setupTest(t)

			ctx := context.Background()
			req := application.SignUpRequest{
				Username: "test",
				Email:    "test@gmail.com",
				Password: "test",
				Fullname: "test",
				Status:   "test",
			}

			mocks.transactor.EXPECT().
				PerformTX(ctx, mock.Anything).
				RunAndReturn(
					func(ctx context.Context, fn func(context.Context, domain.Entities) error) error {
						return fn(ctx, domain.Entities{
							User: mocks.userEntity,
						})
					},
				)

			userEntityCreateReq := entity.UserCreateRequest{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
				Fullname: req.Fullname,
				Status:   req.Status,
			}
			expeted := model.User{
				ID:             "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				Username:       req.Username,
				Email:          req.Email,
				HashedPassword: "test_hashed",
				Fullname:       req.Fullname,
				Status:         req.Status,
				CreatedAt:      time.Now(),
			}
			mocks.passwordMaker.EXPECT().HashPassword(req.Password).Return("test_hashed", nil)
			mocks.userEntity.EXPECT().Create(ctx, userEntityCreateReq).Return(expeted, nil)

			sendEmail := service.Email{
				Subject: "Verification Email for Netsly!",
				Content: fmt.Sprintf("%s, congrats with joining the Netsly community!", req.Username),
				To:      []string{req.Email},
			}
			mocks.taskQueue.EXPECT().SendEmailTask(ctx, service.Critical, sendEmail).Return(nil)

			actual, err := signUsecase.SignUp(ctx, req)
			r.NoError(err)
			r.Equal(expeted, actual)
		})

		t.Run("SignIn", func(t *testing.T) {
			r, mocks, signUsecase := setupTest(t)

			ctx := context.Background()
			req := application.SignInRequest{
				Email:    "test@gmail.com",
				Password: "test-test",
			}
			expected := model.User{
				ID:             "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				HashedPassword: "test-test_hashed",
			}

			mocks.userQuery.EXPECT().GetByEmail(ctx, req.Email).Return(expected, nil)

			mocks.passwordMaker.EXPECT().CheckPassword(req.Password, expected.HashedPassword).Return(nil)

			actual, err := signUsecase.SignIn(ctx, req)
			r.NoError(err)
			r.Equal(expected, actual)
		})
	})
}
