// Code generated by mockery. DO NOT EDIT.

package mockCommand

import (
	context "context"

	command "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"

	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

// Room is an autogenerated mock type for the IRoom type
type Room struct {
	mock.Mock
}

type Room_Expecter struct {
	mock *mock.Mock
}

func (_m *Room) EXPECT() *Room_Expecter {
	return &Room_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *Room) Create(ctx context.Context, req command.RoomCreateRequest) (model.Room, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, command.RoomCreateRequest) (model.Room, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, command.RoomCreateRequest) model.Room); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, command.RoomCreateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Room_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req command.RoomCreateRequest
func (_e *Room_Expecter) Create(ctx interface{}, req interface{}) *Room_Create_Call {
	return &Room_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *Room_Create_Call) Run(run func(ctx context.Context, req command.RoomCreateRequest)) *Room_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(command.RoomCreateRequest))
	})
	return _c
}

func (_c *Room_Create_Call) Return(_a0 model.Room, _a1 error) *Room_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Create_Call) RunAndReturn(run func(context.Context, command.RoomCreateRequest) (model.Room, error)) *Room_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Room) Delete(ctx context.Context, id string) error {
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

// Room_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Room_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Room_Expecter) Delete(ctx interface{}, id interface{}) *Room_Delete_Call {
	return &Room_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *Room_Delete_Call) Run(run func(ctx context.Context, id string)) *Room_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Room_Delete_Call) Return(_a0 error) *Room_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Room_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Room_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, req
func (_m *Room) Update(ctx context.Context, id string, req command.RoomUpdateRequest) (model.Room, error) {
	ret := _m.Called(ctx, id, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, command.RoomUpdateRequest) (model.Room, error)); ok {
		return rf(ctx, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, command.RoomUpdateRequest) model.Room); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, command.RoomUpdateRequest) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Room_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - req command.RoomUpdateRequest
func (_e *Room_Expecter) Update(ctx interface{}, id interface{}, req interface{}) *Room_Update_Call {
	return &Room_Update_Call{Call: _e.mock.On("Update", ctx, id, req)}
}

func (_c *Room_Update_Call) Run(run func(ctx context.Context, id string, req command.RoomUpdateRequest)) *Room_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(command.RoomUpdateRequest))
	})
	return _c
}

func (_c *Room_Update_Call) Return(_a0 model.Room, _a1 error) *Room_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Update_Call) RunAndReturn(run func(context.Context, string, command.RoomUpdateRequest) (model.Room, error)) *Room_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoom creates a new instance of Room. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoom(t interface {
	mock.TestingT
	Cleanup(func())
}) *Room {
	mock := &Room{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
