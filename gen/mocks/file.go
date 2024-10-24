// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/abc-valera/netsly-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-golang/internal/domain/model"
)

// File is an autogenerated mock type for the IFile type
type File struct {
	mock.Mock
}

type File_Expecter struct {
	mock *mock.Mock
}

func (_m *File) EXPECT() *File_Expecter {
	return &File_Expecter{mock: &_m.Mock}
}

// CreateForJoke provides a mock function with given fields: ctx, jokeID, req
func (_m *File) CreateForJoke(ctx context.Context, jokeID string, req entity.FileCreateRequest) (entity.FileCreateResponse, error) {
	ret := _m.Called(ctx, jokeID, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateForJoke")
	}

	var r0 entity.FileCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileCreateRequest) (entity.FileCreateResponse, error)); ok {
		return rf(ctx, jokeID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileCreateRequest) entity.FileCreateResponse); ok {
		r0 = rf(ctx, jokeID, req)
	} else {
		r0 = ret.Get(0).(entity.FileCreateResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.FileCreateRequest) error); ok {
		r1 = rf(ctx, jokeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File_CreateForJoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateForJoke'
type File_CreateForJoke_Call struct {
	*mock.Call
}

// CreateForJoke is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
//   - req entity.FileCreateRequest
func (_e *File_Expecter) CreateForJoke(ctx interface{}, jokeID interface{}, req interface{}) *File_CreateForJoke_Call {
	return &File_CreateForJoke_Call{Call: _e.mock.On("CreateForJoke", ctx, jokeID, req)}
}

func (_c *File_CreateForJoke_Call) Run(run func(ctx context.Context, jokeID string, req entity.FileCreateRequest)) *File_CreateForJoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.FileCreateRequest))
	})
	return _c
}

func (_c *File_CreateForJoke_Call) Return(_a0 entity.FileCreateResponse, _a1 error) *File_CreateForJoke_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *File_CreateForJoke_Call) RunAndReturn(run func(context.Context, string, entity.FileCreateRequest) (entity.FileCreateResponse, error)) *File_CreateForJoke_Call {
	_c.Call.Return(run)
	return _c
}

// CreateForRoom provides a mock function with given fields: ctx, roomID, req
func (_m *File) CreateForRoom(ctx context.Context, roomID string, req entity.FileCreateRequest) (entity.FileCreateResponse, error) {
	ret := _m.Called(ctx, roomID, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateForRoom")
	}

	var r0 entity.FileCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileCreateRequest) (entity.FileCreateResponse, error)); ok {
		return rf(ctx, roomID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileCreateRequest) entity.FileCreateResponse); ok {
		r0 = rf(ctx, roomID, req)
	} else {
		r0 = ret.Get(0).(entity.FileCreateResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.FileCreateRequest) error); ok {
		r1 = rf(ctx, roomID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File_CreateForRoom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateForRoom'
type File_CreateForRoom_Call struct {
	*mock.Call
}

// CreateForRoom is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
//   - req entity.FileCreateRequest
func (_e *File_Expecter) CreateForRoom(ctx interface{}, roomID interface{}, req interface{}) *File_CreateForRoom_Call {
	return &File_CreateForRoom_Call{Call: _e.mock.On("CreateForRoom", ctx, roomID, req)}
}

func (_c *File_CreateForRoom_Call) Run(run func(ctx context.Context, roomID string, req entity.FileCreateRequest)) *File_CreateForRoom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.FileCreateRequest))
	})
	return _c
}

func (_c *File_CreateForRoom_Call) Return(_a0 entity.FileCreateResponse, _a1 error) *File_CreateForRoom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *File_CreateForRoom_Call) RunAndReturn(run func(context.Context, string, entity.FileCreateRequest) (entity.FileCreateResponse, error)) *File_CreateForRoom_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *File) Delete(ctx context.Context, id string) error {
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

// File_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type File_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *File_Expecter) Delete(ctx interface{}, id interface{}) *File_Delete_Call {
	return &File_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *File_Delete_Call) Run(run func(ctx context.Context, id string)) *File_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *File_Delete_Call) Return(_a0 error) *File_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Delete_Call) RunAndReturn(run func(context.Context, string) error) *File_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *File) GetByID(ctx context.Context, id string) (model.FileInfo, model.FileContent, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.FileInfo
	var r1 model.FileContent
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FileInfo, model.FileContent, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FileInfo); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.FileInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) model.FileContent); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(model.FileContent)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// File_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type File_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *File_Expecter) GetByID(ctx interface{}, id interface{}) *File_GetByID_Call {
	return &File_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *File_GetByID_Call) Run(run func(ctx context.Context, id string)) *File_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *File_GetByID_Call) Return(_a0 model.FileInfo, _a1 model.FileContent, _a2 error) *File_GetByID_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *File_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.FileInfo, model.FileContent, error)) *File_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, req
func (_m *File) Update(ctx context.Context, id string, req entity.FileUpdateRequest) (model.FileInfo, error) {
	ret := _m.Called(ctx, id, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileUpdateRequest) (model.FileInfo, error)); ok {
		return rf(ctx, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.FileUpdateRequest) model.FileInfo); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Get(0).(model.FileInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.FileUpdateRequest) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type File_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - req entity.FileUpdateRequest
func (_e *File_Expecter) Update(ctx interface{}, id interface{}, req interface{}) *File_Update_Call {
	return &File_Update_Call{Call: _e.mock.On("Update", ctx, id, req)}
}

func (_c *File_Update_Call) Run(run func(ctx context.Context, id string, req entity.FileUpdateRequest)) *File_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.FileUpdateRequest))
	})
	return _c
}

func (_c *File_Update_Call) Return(_a0 model.FileInfo, _a1 error) *File_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *File_Update_Call) RunAndReturn(run func(context.Context, string, entity.FileUpdateRequest) (model.FileInfo, error)) *File_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewFile creates a new instance of File. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *File {
	mock := &File{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
