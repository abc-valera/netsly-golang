// Code generated by mockery. DO NOT EDIT.

package mockEntity

import (
	context "context"

	entity "github.com/abc-valera/netsly-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-golang/internal/domain/model"

	selector "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
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
func (_m *Room) Create(ctx context.Context, req entity.RoomCreateRequest) (model.Room, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomCreateRequest) (model.Room, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomCreateRequest) model.Room); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.RoomCreateRequest) error); ok {
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
//   - req entity.RoomCreateRequest
func (_e *Room_Expecter) Create(ctx interface{}, req interface{}) *Room_Create_Call {
	return &Room_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *Room_Create_Call) Run(run func(ctx context.Context, req entity.RoomCreateRequest)) *Room_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.RoomCreateRequest))
	})
	return _c
}

func (_c *Room_Create_Call) Return(_a0 model.Room, _a1 error) *Room_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Create_Call) RunAndReturn(run func(context.Context, entity.RoomCreateRequest) (model.Room, error)) *Room_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, roomID
func (_m *Room) Delete(ctx context.Context, roomID string) error {
	ret := _m.Called(ctx, roomID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, roomID)
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
//   - roomID string
func (_e *Room_Expecter) Delete(ctx interface{}, roomID interface{}) *Room_Delete_Call {
	return &Room_Delete_Call{Call: _e.mock.On("Delete", ctx, roomID)}
}

func (_c *Room_Delete_Call) Run(run func(ctx context.Context, roomID string)) *Room_Delete_Call {
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

// GetAllByUserID provides a mock function with given fields: ctx, userID, s
func (_m *Room) GetAllByUserID(ctx context.Context, userID string, s selector.Selector) (model.Rooms, error) {
	ret := _m.Called(ctx, userID, s)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByUserID")
	}

	var r0 model.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Rooms, error)); ok {
		return rf(ctx, userID, s)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Rooms); ok {
		r0 = rf(ctx, userID, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, userID, s)
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
//   - s selector.Selector
func (_e *Room_Expecter) GetAllByUserID(ctx interface{}, userID interface{}, s interface{}) *Room_GetAllByUserID_Call {
	return &Room_GetAllByUserID_Call{Call: _e.mock.On("GetAllByUserID", ctx, userID, s)}
}

func (_c *Room_GetAllByUserID_Call) Run(run func(ctx context.Context, userID string, s selector.Selector)) *Room_GetAllByUserID_Call {
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

// SearchAllByName provides a mock function with given fields: ctx, keyword, s
func (_m *Room) SearchAllByName(ctx context.Context, keyword string, s selector.Selector) (model.Rooms, error) {
	ret := _m.Called(ctx, keyword, s)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByName")
	}

	var r0 model.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Rooms, error)); ok {
		return rf(ctx, keyword, s)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Rooms); ok {
		r0 = rf(ctx, keyword, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, keyword, s)
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
//   - s selector.Selector
func (_e *Room_Expecter) SearchAllByName(ctx interface{}, keyword interface{}, s interface{}) *Room_SearchAllByName_Call {
	return &Room_SearchAllByName_Call{Call: _e.mock.On("SearchAllByName", ctx, keyword, s)}
}

func (_c *Room_SearchAllByName_Call) Run(run func(ctx context.Context, keyword string, s selector.Selector)) *Room_SearchAllByName_Call {
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

// Update provides a mock function with given fields: ctx, roomID, req
func (_m *Room) Update(ctx context.Context, roomID string, req entity.RoomUpdateRequest) (model.Room, error) {
	ret := _m.Called(ctx, roomID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomUpdateRequest) (model.Room, error)); ok {
		return rf(ctx, roomID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomUpdateRequest) model.Room); ok {
		r0 = rf(ctx, roomID, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.RoomUpdateRequest) error); ok {
		r1 = rf(ctx, roomID, req)
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
//   - roomID string
//   - req entity.RoomUpdateRequest
func (_e *Room_Expecter) Update(ctx interface{}, roomID interface{}, req interface{}) *Room_Update_Call {
	return &Room_Update_Call{Call: _e.mock.On("Update", ctx, roomID, req)}
}

func (_c *Room_Update_Call) Run(run func(ctx context.Context, roomID string, req entity.RoomUpdateRequest)) *Room_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.RoomUpdateRequest))
	})
	return _c
}

func (_c *Room_Update_Call) Return(_a0 model.Room, _a1 error) *Room_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Update_Call) RunAndReturn(run func(context.Context, string, entity.RoomUpdateRequest) (model.Room, error)) *Room_Update_Call {
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
