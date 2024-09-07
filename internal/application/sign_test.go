package application

import (
	"testing"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/google/uuid"
)

func TestSignUsecase(t *testing.T) {
	t.Run("SignUp", func(t *testing.T) {
		defaultCtx, instrumentedCtx, r, dep := NewTest(t)

		req := SignUpRequest{
			Username: "username",
			Email:    "email@gmail.com",
			Password: "test_password",
			Fullname: "fullname",
			Status:   "status",
		}

		expeted := model.User{
			ID:             uuid.NewString(),
			Username:       req.Username,
			Email:          req.Email,
			HashedPassword: "test_password_hashed",
			Fullname:       req.Fullname,
			Status:         req.Status,
			CreatedAt:      time.Now(),
		}

		dep.MockEntities.User.EXPECT().Create(instrumentedCtx, entity.UserCreateRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Fullname: req.Fullname,
			Status:   req.Status,
		}).Return(expeted, nil)

		dep.MockEntities.Emailer.EXPECT().
			SendEmail(welcomeEmailTemplateFunc(req.Username, req.Email)).
			Return(nil)

		actual, err := newSignUsecase(dep).SignUp(defaultCtx, req)
		r.NoError(err)
		r.Equal(expeted, actual)
	})

	t.Run("SignIn", func(t *testing.T) {
		defaultContext, instrumentedContext, r, dep := NewTest(t)

		req := SignInRequest{
			Email:    "test@gmail.com",
			Password: "test-test",
		}
		expected := model.User{
			ID:             uuid.NewString(),
			HashedPassword: "test_password_hashed",
		}

		dep.MockEntities.User.EXPECT().GetByEmail(instrumentedContext, req.Email).Return(expected, nil)

		dep.MockEntities.Passworder.EXPECT().CheckPassword(req.Password, expected.HashedPassword).Return(nil)

		actual, err := newSignUsecase(dep).SignIn(defaultContext, req)
		r.NoError(err)
		r.Equal(expected, actual)
	})
}
