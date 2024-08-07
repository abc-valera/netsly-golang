// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	selector "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
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

// GetAllByUserID provides a mock function with given fields: ctx, userID, _a2
func (_m *Room) GetAllByUserID(ctx context.Context, userID string, _a2 selector.Selector) (model.Rooms, error) {
	ret := _m.Called(ctx, userID, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByUserID")
	}

	var r0 model.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Rooms, error)); ok {
		return rf(ctx, userID, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Rooms); ok {
		r0 = rf(ctx, userID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, userID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_GetAllByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUserID'
type Room_GetAllByUserID_Call struct {
	*mock.Call
}

// GetAllByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - _a2 selector.Selector
func (_e *Room_Expecter) GetAllByUserID(ctx interface{}, userID interface{}, _a2 interface{}) *Room_GetAllByUserID_Call {
	return &Room_GetAllByUserID_Call{Call: _e.mock.On("GetAllByUserID", ctx, userID, _a2)}
}

func (_c *Room_GetAllByUserID_Call) Run(run func(ctx context.Context, userID string, _a2 selector.Selector)) *Room_GetAllByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Room_GetAllByUserID_Call) Return(_a0 model.Rooms, _a1 error) *Room_GetAllByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_GetAllByUserID_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Rooms, error)) *Room_GetAllByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Room) GetByID(ctx context.Context, id string) (model.Room, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Room, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Room); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type Room_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Room_Expecter) GetByID(ctx interface{}, id interface{}) *Room_GetByID_Call {
	return &Room_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *Room_GetByID_Call) Run(run func(ctx context.Context, id string)) *Room_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Room_GetByID_Call) Return(_a0 model.Room, _a1 error) *Room_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.Room, error)) *Room_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *Room) GetByName(ctx context.Context, name string) (model.Room, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Room, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Room); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type Room_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *Room_Expecter) GetByName(ctx interface{}, name interface{}) *Room_GetByName_Call {
	return &Room_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *Room_GetByName_Call) Run(run func(ctx context.Context, name string)) *Room_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Room_GetByName_Call) Return(_a0 model.Room, _a1 error) *Room_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_GetByName_Call) RunAndReturn(run func(context.Context, string) (model.Room, error)) *Room_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAllByName provides a mock function with given fields: ctx, keyword, _a2
func (_m *Room) SearchAllByName(ctx context.Context, keyword string, _a2 selector.Selector) (model.Rooms, error) {
	ret := _m.Called(ctx, keyword, _a2)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByName")
	}

	var r0 model.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Rooms, error)); ok {
		return rf(ctx, keyword, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Rooms); ok {
		r0 = rf(ctx, keyword, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, keyword, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_SearchAllByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAllByName'
type Room_SearchAllByName_Call struct {
	*mock.Call
}

// SearchAllByName is a helper method to define mock.On call
//   - ctx context.Context
//   - keyword string
//   - _a2 selector.Selector
func (_e *Room_Expecter) SearchAllByName(ctx interface{}, keyword interface{}, _a2 interface{}) *Room_SearchAllByName_Call {
	return &Room_SearchAllByName_Call{Call: _e.mock.On("SearchAllByName", ctx, keyword, _a2)}
}

func (_c *Room_SearchAllByName_Call) Run(run func(ctx context.Context, keyword string, _a2 selector.Selector)) *Room_SearchAllByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Room_SearchAllByName_Call) Return(_a0 model.Rooms, _a1 error) *Room_SearchAllByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_SearchAllByName_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Rooms, error)) *Room_SearchAllByName_Call {
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
