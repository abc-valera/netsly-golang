// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	mock "github.com/stretchr/testify/mock"
)

// RoomMember is an autogenerated mock type for the IRoomMember type
type RoomMember struct {
	mock.Mock
}

type RoomMember_Expecter struct {
	mock *mock.Mock
}

func (_m *RoomMember) EXPECT() *RoomMember_Expecter {
	return &RoomMember_Expecter{mock: &_m.Mock}
}

// GetByIDs provides a mock function with given fields: ctx, userID, roomID
func (_m *RoomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	ret := _m.Called(ctx, userID, roomID)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 model.RoomMember
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (model.RoomMember, error)); ok {
		return rf(ctx, userID, roomID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) model.RoomMember); ok {
		r0 = rf(ctx, userID, roomID)
	} else {
		r0 = ret.Get(0).(model.RoomMember)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMember_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type RoomMember_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - roomID string
func (_e *RoomMember_Expecter) GetByIDs(ctx interface{}, userID interface{}, roomID interface{}) *RoomMember_GetByIDs_Call {
	return &RoomMember_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, userID, roomID)}
}

func (_c *RoomMember_GetByIDs_Call) Run(run func(ctx context.Context, userID string, roomID string)) *RoomMember_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *RoomMember_GetByIDs_Call) Return(_a0 model.RoomMember, _a1 error) *RoomMember_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMember_GetByIDs_Call) RunAndReturn(run func(context.Context, string, string) (model.RoomMember, error)) *RoomMember_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetByRoomID provides a mock function with given fields: ctx, roomID
func (_m *RoomMember) GetByRoomID(ctx context.Context, roomID string) (model.RoomMembers, error) {
	ret := _m.Called(ctx, roomID)

	if len(ret) == 0 {
		panic("no return value specified for GetByRoomID")
	}

	var r0 model.RoomMembers
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.RoomMembers, error)); ok {
		return rf(ctx, roomID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.RoomMembers); ok {
		r0 = rf(ctx, roomID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.RoomMembers)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMember_GetByRoomID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByRoomID'
type RoomMember_GetByRoomID_Call struct {
	*mock.Call
}

// GetByRoomID is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
func (_e *RoomMember_Expecter) GetByRoomID(ctx interface{}, roomID interface{}) *RoomMember_GetByRoomID_Call {
	return &RoomMember_GetByRoomID_Call{Call: _e.mock.On("GetByRoomID", ctx, roomID)}
}

func (_c *RoomMember_GetByRoomID_Call) Run(run func(ctx context.Context, roomID string)) *RoomMember_GetByRoomID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RoomMember_GetByRoomID_Call) Return(_a0 model.RoomMembers, _a1 error) *RoomMember_GetByRoomID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMember_GetByRoomID_Call) RunAndReturn(run func(context.Context, string) (model.RoomMembers, error)) *RoomMember_GetByRoomID_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoomMember creates a new instance of RoomMember. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoomMember(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoomMember {
	mock := &RoomMember{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}