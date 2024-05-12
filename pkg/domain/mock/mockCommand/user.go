// Code generated by mockery. DO NOT EDIT.

package mockCommand

import (
	context "context"

	command "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"

	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-api-golang/pkg/domain/model"
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

// Create provides a mock function with given fields: ctx, req
func (_m *User) Create(ctx context.Context, req model.User) (model.User, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) (model.User, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User) model.User); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type User_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req model.User
func (_e *User_Expecter) Create(ctx interface{}, req interface{}) *User_Create_Call {
	return &User_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *User_Create_Call) Run(run func(ctx context.Context, req model.User)) *User_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.User))
	})
	return _c
}

func (_c *User_Create_Call) Return(_a0 model.User, _a1 error) *User_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_Create_Call) RunAndReturn(run func(context.Context, model.User) (model.User, error)) *User_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *User) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// User_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type User_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *User_Expecter) Delete(ctx interface{}, id interface{}) *User_Delete_Call {
	return &User_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *User_Delete_Call) Run(run func(ctx context.Context, id string)) *User_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *User_Delete_Call) Return(_a0 error) *User_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *User_Delete_Call) RunAndReturn(run func(context.Context, string) error) *User_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, req
func (_m *User) Update(ctx context.Context, id string, req command.UserUpdate) (model.User, error) {
	ret := _m.Called(ctx, id, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, command.UserUpdate) (model.User, error)); ok {
		return rf(ctx, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, command.UserUpdate) model.User); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, command.UserUpdate) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type User_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - req command.UserUpdate
func (_e *User_Expecter) Update(ctx interface{}, id interface{}, req interface{}) *User_Update_Call {
	return &User_Update_Call{Call: _e.mock.On("Update", ctx, id, req)}
}

func (_c *User_Update_Call) Run(run func(ctx context.Context, id string, req command.UserUpdate)) *User_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(command.UserUpdate))
	})
	return _c
}

func (_c *User_Update_Call) Return(_a0 model.User, _a1 error) *User_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *User_Update_Call) RunAndReturn(run func(context.Context, string, command.UserUpdate) (model.User, error)) *User_Update_Call {
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
