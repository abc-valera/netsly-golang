package application

import (
	"context"
	"testing"
	"time"

	"github.com/abc-valera/netsly-api-golang/gen/mock/mockEntity"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockEntityTransactor"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockPassworder"
	"github.com/abc-valera/netsly-api-golang/gen/mock/mockTaskQueuer"
	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/nopLogger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestSignUsecase(t *testing.T) {
	type Mocks struct {
		userEntity *mockEntity.User
		passworder *mockPassworder.Passworder
		taskQueue  *mockTaskQueuer.TaskQueuer
		transactor *mockEntityTransactor.Transactor
	}

	// Init global variables
	// (generally we don't want to mock it, just use noop variants and make sure it's not null)
	global.Init(
		mode.Production,
		noop.NewTracerProvider().Tracer("noop"),
		nopLogger.New(),
	)

	// setupTest is a helper function to setup the test.
	// It returns the instrumented background context, instance of require.Assertions,
	// initialized mocks, and the usecase, initialized with the mocks.
	setupTest := func(t *testing.T) (context.Context, *require.Assertions, Mocks, ISignUsecase) {
		mocks := Mocks{
			userEntity: mockEntity.NewUser(t),
			passworder: mockPassworder.NewPassworder(t),
			taskQueue:  mockTaskQueuer.NewTaskQueuer(t),
			transactor: mockEntityTransactor.NewTransactor(t),
		}

		ctx, _ := global.NewSpan(context.Background())

		return ctx, require.New(t), mocks, NewSignUsecase(
			mocks.userEntity,
			mocks.transactor,
			mocks.passworder,
			mocks.taskQueue,
		)
	}

	t.Run("SignUsecase", func(t *testing.T) {
		t.Run("SignUp", func(t *testing.T) {
			ctx, r, mocks, signUsecase := setupTest(t)

			req := SignUpRequest{
				Username: "username",
				Email:    "email@gmail.com",
				Password: "test_password",
				Fullname: "fullname",
				Status:   "status",
			}

			expetedResp := model.User{
				ID:             uuid.NewString(),
				Username:       req.Username,
				Email:          req.Email,
				HashedPassword: "test_password_hashed",
				Fullname:       req.Fullname,
				Status:         req.Status,
				CreatedAt:      time.Now(),
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

			mocks.userEntity.EXPECT().Create(ctx, userEntityCreateReq).Return(expetedResp, nil)

			mocks.taskQueue.EXPECT().SendEmailTask(
				ctx,
				service.Critical,
				welcomeEmailTemplateFunc(req.Username, req.Email),
			).Return(nil)

			// Note, that the usecase should be ran with newly created background context
			// (setupTest function returns the already instrumented one of it)
			actualResp, err := signUsecase.SignUp(context.Background(), req)
			r.NoError(err)
			r.Equal(expetedResp, actualResp)
		})

		t.Run("SignIn", func(t *testing.T) {
			ctx, r, mocks, signUsecase := setupTest(t)

			req := SignInRequest{
				Email:    "test@gmail.com",
				Password: "test-test",
			}
			expected := model.User{
				ID:             uuid.NewString(),
				HashedPassword: "test_password_hashed",
			}

			mocks.userEntity.EXPECT().GetByEmail(ctx, req.Email).Return(expected, nil)

			mocks.passworder.EXPECT().CheckPassword(req.Password, expected.HashedPassword).Return(nil)

			actual, err := signUsecase.SignIn(context.Background(), req)
			r.NoError(err)
			r.Equal(expected, actual)
		})
	})
}
