// Code generated by mockery. DO NOT EDIT.

package mockEntity

import (
	context "context"

	entity "github.com/abc-valera/netsly-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-golang/internal/domain/model"

	selector "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

// RoomMessage is an autogenerated mock type for the IRoomMessage type
type RoomMessage struct {
	mock.Mock
}

type RoomMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *RoomMessage) EXPECT() *RoomMessage_Expecter {
	return &RoomMessage_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *RoomMessage) Create(ctx context.Context, req entity.RoomMessageCreateRequest) (model.RoomMessage, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.RoomMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomMessageCreateRequest) (model.RoomMessage, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomMessageCreateRequest) model.RoomMessage); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.RoomMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.RoomMessageCreateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMessage_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type RoomMessage_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.RoomMessageCreateRequest
func (_e *RoomMessage_Expecter) Create(ctx interface{}, req interface{}) *RoomMessage_Create_Call {
	return &RoomMessage_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *RoomMessage_Create_Call) Run(run func(ctx context.Context, req entity.RoomMessageCreateRequest)) *RoomMessage_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.RoomMessageCreateRequest))
	})
	return _c
}

func (_c *RoomMessage_Create_Call) Return(_a0 model.RoomMessage, _a1 error) *RoomMessage_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMessage_Create_Call) RunAndReturn(run func(context.Context, entity.RoomMessageCreateRequest) (model.RoomMessage, error)) *RoomMessage_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *RoomMessage) Delete(ctx context.Context, id string) error {
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

// RoomMessage_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type RoomMessage_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *RoomMessage_Expecter) Delete(ctx interface{}, id interface{}) *RoomMessage_Delete_Call {
	return &RoomMessage_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *RoomMessage_Delete_Call) Run(run func(ctx context.Context, id string)) *RoomMessage_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RoomMessage_Delete_Call) Return(_a0 error) *RoomMessage_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoomMessage_Delete_Call) RunAndReturn(run func(context.Context, string) error) *RoomMessage_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByRoomID provides a mock function with given fields: ctx, roomID, s
func (_m *RoomMessage) GetAllByRoomID(ctx context.Context, roomID string, s selector.Selector) (model.RoomMessages, error) {
	ret := _m.Called(ctx, roomID, s)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByRoomID")
	}

	var r0 model.RoomMessages
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.RoomMessages, error)); ok {
		return rf(ctx, roomID, s)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.RoomMessages); ok {
		r0 = rf(ctx, roomID, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.RoomMessages)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, roomID, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMessage_GetAllByRoomID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByRoomID'
type RoomMessage_GetAllByRoomID_Call struct {
	*mock.Call
}

// GetAllByRoomID is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
//   - s selector.Selector
func (_e *RoomMessage_Expecter) GetAllByRoomID(ctx interface{}, roomID interface{}, s interface{}) *RoomMessage_GetAllByRoomID_Call {
	return &RoomMessage_GetAllByRoomID_Call{Call: _e.mock.On("GetAllByRoomID", ctx, roomID, s)}
}

func (_c *RoomMessage_GetAllByRoomID_Call) Run(run func(ctx context.Context, roomID string, s selector.Selector)) *RoomMessage_GetAllByRoomID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *RoomMessage_GetAllByRoomID_Call) Return(_a0 model.RoomMessages, _a1 error) *RoomMessage_GetAllByRoomID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMessage_GetAllByRoomID_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.RoomMessages, error)) *RoomMessage_GetAllByRoomID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *RoomMessage) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.RoomMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.RoomMessage, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.RoomMessage); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.RoomMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMessage_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type RoomMessage_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *RoomMessage_Expecter) GetByID(ctx interface{}, id interface{}) *RoomMessage_GetByID_Call {
	return &RoomMessage_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *RoomMessage_GetByID_Call) Run(run func(ctx context.Context, id string)) *RoomMessage_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RoomMessage_GetByID_Call) Return(_a0 model.RoomMessage, _a1 error) *RoomMessage_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMessage_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.RoomMessage, error)) *RoomMessage_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAllByText provides a mock function with given fields: ctx, keyword, s
func (_m *RoomMessage) SearchAllByText(ctx context.Context, keyword string, s selector.Selector) (model.RoomMessages, error) {
	ret := _m.Called(ctx, keyword, s)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByText")
	}

	var r0 model.RoomMessages
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.RoomMessages, error)); ok {
		return rf(ctx, keyword, s)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.RoomMessages); ok {
		r0 = rf(ctx, keyword, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.RoomMessages)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, keyword, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMessage_SearchAllByText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAllByText'
type RoomMessage_SearchAllByText_Call struct {
	*mock.Call
}

// SearchAllByText is a helper method to define mock.On call
//   - ctx context.Context
//   - keyword string
//   - s selector.Selector
func (_e *RoomMessage_Expecter) SearchAllByText(ctx interface{}, keyword interface{}, s interface{}) *RoomMessage_SearchAllByText_Call {
	return &RoomMessage_SearchAllByText_Call{Call: _e.mock.On("SearchAllByText", ctx, keyword, s)}
}

func (_c *RoomMessage_SearchAllByText_Call) Run(run func(ctx context.Context, keyword string, s selector.Selector)) *RoomMessage_SearchAllByText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *RoomMessage_SearchAllByText_Call) Return(_a0 model.RoomMessages, _a1 error) *RoomMessage_SearchAllByText_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMessage_SearchAllByText_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.RoomMessages, error)) *RoomMessage_SearchAllByText_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, req
func (_m *RoomMessage) Update(ctx context.Context, id string, req entity.RoomMessageUpdateRequest) (model.RoomMessage, error) {
	ret := _m.Called(ctx, id, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.RoomMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomMessageUpdateRequest) (model.RoomMessage, error)); ok {
		return rf(ctx, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomMessageUpdateRequest) model.RoomMessage); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Get(0).(model.RoomMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.RoomMessageUpdateRequest) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoomMessage_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type RoomMessage_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - req entity.RoomMessageUpdateRequest
func (_e *RoomMessage_Expecter) Update(ctx interface{}, id interface{}, req interface{}) *RoomMessage_Update_Call {
	return &RoomMessage_Update_Call{Call: _e.mock.On("Update", ctx, id, req)}
}

func (_c *RoomMessage_Update_Call) Run(run func(ctx context.Context, id string, req entity.RoomMessageUpdateRequest)) *RoomMessage_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.RoomMessageUpdateRequest))
	})
	return _c
}

func (_c *RoomMessage_Update_Call) Return(_a0 model.RoomMessage, _a1 error) *RoomMessage_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoomMessage_Update_Call) RunAndReturn(run func(context.Context, string, entity.RoomMessageUpdateRequest) (model.RoomMessage, error)) *RoomMessage_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoomMessage creates a new instance of RoomMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoomMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoomMessage {
	mock := &RoomMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
