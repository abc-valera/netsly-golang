// Code generated by mockery. DO NOT EDIT.

package mockTokenMaker

import (
	service "github.com/abc-valera/netsly-api-golang/pkg/domain/service"
	mock "github.com/stretchr/testify/mock"
)

// TokenMaker is an autogenerated mock type for the ITokenMaker type
type TokenMaker struct {
	mock.Mock
}

type TokenMaker_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenMaker) EXPECT() *TokenMaker_Expecter {
	return &TokenMaker_Expecter{mock: &_m.Mock}
}

// CreateAccessToken provides a mock function with given fields: userID
func (_m *TokenMaker) CreateAccessToken(userID string) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenMaker_CreateAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccessToken'
type TokenMaker_CreateAccessToken_Call struct {
	*mock.Call
}

// CreateAccessToken is a helper method to define mock.On call
//   - userID string
func (_e *TokenMaker_Expecter) CreateAccessToken(userID interface{}) *TokenMaker_CreateAccessToken_Call {
	return &TokenMaker_CreateAccessToken_Call{Call: _e.mock.On("CreateAccessToken", userID)}
}

func (_c *TokenMaker_CreateAccessToken_Call) Run(run func(userID string)) *TokenMaker_CreateAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TokenMaker_CreateAccessToken_Call) Return(_a0 string, _a1 error) *TokenMaker_CreateAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenMaker_CreateAccessToken_Call) RunAndReturn(run func(string) (string, error)) *TokenMaker_CreateAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRefreshToken provides a mock function with given fields: userID
func (_m *TokenMaker) CreateRefreshToken(userID string) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenMaker_CreateRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRefreshToken'
type TokenMaker_CreateRefreshToken_Call struct {
	*mock.Call
}

// CreateRefreshToken is a helper method to define mock.On call
//   - userID string
func (_e *TokenMaker_Expecter) CreateRefreshToken(userID interface{}) *TokenMaker_CreateRefreshToken_Call {
	return &TokenMaker_CreateRefreshToken_Call{Call: _e.mock.On("CreateRefreshToken", userID)}
}

func (_c *TokenMaker_CreateRefreshToken_Call) Run(run func(userID string)) *TokenMaker_CreateRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TokenMaker_CreateRefreshToken_Call) Return(_a0 string, _a1 error) *TokenMaker_CreateRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenMaker_CreateRefreshToken_Call) RunAndReturn(run func(string) (string, error)) *TokenMaker_CreateRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyToken provides a mock function with given fields: token
func (_m *TokenMaker) VerifyToken(token string) (service.AuthPayload, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 service.AuthPayload
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (service.AuthPayload, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) service.AuthPayload); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(service.AuthPayload)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenMaker_VerifyToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyToken'
type TokenMaker_VerifyToken_Call struct {
	*mock.Call
}

// VerifyToken is a helper method to define mock.On call
//   - token string
func (_e *TokenMaker_Expecter) VerifyToken(token interface{}) *TokenMaker_VerifyToken_Call {
	return &TokenMaker_VerifyToken_Call{Call: _e.mock.On("VerifyToken", token)}
}

func (_c *TokenMaker_VerifyToken_Call) Run(run func(token string)) *TokenMaker_VerifyToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TokenMaker_VerifyToken_Call) Return(_a0 service.AuthPayload, _a1 error) *TokenMaker_VerifyToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenMaker_VerifyToken_Call) RunAndReturn(run func(string) (service.AuthPayload, error)) *TokenMaker_VerifyToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenMaker creates a new instance of TokenMaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenMaker(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenMaker {
	mock := &TokenMaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
