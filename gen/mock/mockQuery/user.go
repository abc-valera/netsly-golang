// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-golang/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	selector "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

// User is an autogenerated mock type for the IUser type
type User struct {
	mock.Mock
}

type User_Expecter struct {
	mock *mock.Mock
}

func (_m *User) EXPECT() *User_Expecter {
	return &User_Expecter{mock: &_m.Mock}
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *User) GetByEmail(ctx context.Context, email string) (model.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type User_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *User_Expecter) GetByEmail(ctx interface{}, email interface{}) *User_GetByEmail_Call {
	return &User_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *User_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *User_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *User_GetByEmail_Call) Return(_a0 model.User, _a1 error) *User_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_GetByEmail_Call) RunAndReturn(run func(context.Context, string) (model.User, error)) *User_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *User) GetByID(ctx context.Context, id string) (model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type User_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *User_Expecter) GetByID(ctx interface{}, id interface{}) *User_GetByID_Call {
	return &User_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *User_GetByID_Call) Run(run func(ctx context.Context, id string)) *User_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *User_GetByID_Call) Return(_a0 model.User, _a1 error) *User_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.User, error)) *User_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *User) GetByUsername(ctx context.Context, username string) (model.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_GetByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUsername'
type User_GetByUsername_Call struct {
	*mock.Call
}

// GetByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *User_Expecter) GetByUsername(ctx interface{}, username interface{}) *User_GetByUsername_Call {
	return &User_GetByUsername_Call{Call: _e.mock.On("GetByUsername", ctx, username)}
}

func (_c *User_GetByUsername_Call) Run(run func(ctx context.Context, username string)) *User_GetByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *User_GetByUsername_Call) Return(_a0 model.User, _a1 error) *User_GetByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_GetByUsername_Call) RunAndReturn(run func(context.Context, string) (model.User, error)) *User_GetByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAllByUsername provides a mock function with given fields: ctx, keyword, _a2
func (_m *User) SearchAllByUsername(ctx context.Context, keyword string, _a2 selector.Selector) (model.Users, error) {
	ret := _m.Called(ctx, keyword, _a2)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByUsername")
	}

	var r0 model.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Users, error)); ok {
		return rf(ctx, keyword, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Users); ok {
		r0 = rf(ctx, keyword, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, keyword, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_SearchAllByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAllByUsername'
type User_SearchAllByUsername_Call struct {
	*mock.Call
}

// SearchAllByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - keyword string
//   - _a2 selector.Selector
func (_e *User_Expecter) SearchAllByUsername(ctx interface{}, keyword interface{}, _a2 interface{}) *User_SearchAllByUsername_Call {
	return &User_SearchAllByUsername_Call{Call: _e.mock.On("SearchAllByUsername", ctx, keyword, _a2)}
}

func (_c *User_SearchAllByUsername_Call) Run(run func(ctx context.Context, keyword string, _a2 selector.Selector)) *User_SearchAllByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *User_SearchAllByUsername_Call) Return(_a0 model.Users, _a1 error) *User_SearchAllByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_SearchAllByUsername_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Users, error)) *User_SearchAllByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
