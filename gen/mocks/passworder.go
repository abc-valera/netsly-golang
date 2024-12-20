// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Passworder is an autogenerated mock type for the IPassworder type
type Passworder struct {
	mock.Mock
}

type Passworder_Expecter struct {
	mock *mock.Mock
}

func (_m *Passworder) EXPECT() *Passworder_Expecter {
	return &Passworder_Expecter{mock: &_m.Mock}
}

// CheckPassword provides a mock function with given fields: password, hash
func (_m *Passworder) CheckPassword(password string, hash string) error {
	ret := _m.Called(password, hash)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Passworder_CheckPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckPassword'
type Passworder_CheckPassword_Call struct {
	*mock.Call
}

// CheckPassword is a helper method to define mock.On call
//   - password string
//   - hash string
func (_e *Passworder_Expecter) CheckPassword(password interface{}, hash interface{}) *Passworder_CheckPassword_Call {
	return &Passworder_CheckPassword_Call{Call: _e.mock.On("CheckPassword", password, hash)}
}

func (_c *Passworder_CheckPassword_Call) Run(run func(password string, hash string)) *Passworder_CheckPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Passworder_CheckPassword_Call) Return(_a0 error) *Passworder_CheckPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Passworder_CheckPassword_Call) RunAndReturn(run func(string, string) error) *Passworder_CheckPassword_Call {
	_c.Call.Return(run)
	return _c
}

// HashPassword provides a mock function with given fields: password
func (_m *Passworder) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Passworder_HashPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashPassword'
type Passworder_HashPassword_Call struct {
	*mock.Call
}

// HashPassword is a helper method to define mock.On call
//   - password string
func (_e *Passworder_Expecter) HashPassword(password interface{}) *Passworder_HashPassword_Call {
	return &Passworder_HashPassword_Call{Call: _e.mock.On("HashPassword", password)}
}

func (_c *Passworder_HashPassword_Call) Run(run func(password string)) *Passworder_HashPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Passworder_HashPassword_Call) Return(_a0 string, _a1 error) *Passworder_HashPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Passworder_HashPassword_Call) RunAndReturn(run func(string) (string, error)) *Passworder_HashPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewPassworder creates a new instance of Passworder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPassworder(t interface {
	mock.TestingT
	Cleanup(func())
}) *Passworder {
	mock := &Passworder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
