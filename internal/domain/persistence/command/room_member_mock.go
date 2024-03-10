// Code generated by mockery. DO NOT EDIT.

package command

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	mock "github.com/stretchr/testify/mock"
)

// MockRoomMember is an autogenerated mock type for the IRoomMember type
type MockRoomMember struct {
	mock.Mock
}

type MockRoomMember_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRoomMember) EXPECT() *MockRoomMember_Expecter {
	return &MockRoomMember_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *MockRoomMember) Create(ctx context.Context, req model.RoomMember) (model.RoomMember, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.RoomMember
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RoomMember) (model.RoomMember, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.RoomMember) model.RoomMember); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.RoomMember)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.RoomMember) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRoomMember_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRoomMember_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req model.RoomMember
func (_e *MockRoomMember_Expecter) Create(ctx interface{}, req interface{}) *MockRoomMember_Create_Call {
	return &MockRoomMember_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *MockRoomMember_Create_Call) Run(run func(ctx context.Context, req model.RoomMember)) *MockRoomMember_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.RoomMember))
	})
	return _c
}

func (_c *MockRoomMember_Create_Call) Return(_a0 model.RoomMember, _a1 error) *MockRoomMember_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRoomMember_Create_Call) RunAndReturn(run func(context.Context, model.RoomMember) (model.RoomMember, error)) *MockRoomMember_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, RoomID, UserID
func (_m *MockRoomMember) Delete(ctx context.Context, RoomID string, UserID string) error {
	ret := _m.Called(ctx, RoomID, UserID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, RoomID, UserID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRoomMember_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRoomMember_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - RoomID string
//   - UserID string
func (_e *MockRoomMember_Expecter) Delete(ctx interface{}, RoomID interface{}, UserID interface{}) *MockRoomMember_Delete_Call {
	return &MockRoomMember_Delete_Call{Call: _e.mock.On("Delete", ctx, RoomID, UserID)}
}

func (_c *MockRoomMember_Delete_Call) Run(run func(ctx context.Context, RoomID string, UserID string)) *MockRoomMember_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRoomMember_Delete_Call) Return(_a0 error) *MockRoomMember_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRoomMember_Delete_Call) RunAndReturn(run func(context.Context, string, string) error) *MockRoomMember_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRoomMember creates a new instance of MockRoomMember. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRoomMember(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRoomMember {
	mock := &MockRoomMember{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
