// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	selector "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

// FileInfo is an autogenerated mock type for the IFileInfo type
type FileInfo struct {
	mock.Mock
}

type FileInfo_Expecter struct {
	mock *mock.Mock
}

func (_m *FileInfo) EXPECT() *FileInfo_Expecter {
	return &FileInfo_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields: ctx, _a1
func (_m *FileInfo) GetAll(ctx context.Context, _a1 selector.Selector) (model.FileInfos, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 model.FileInfos
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, selector.Selector) (model.FileInfos, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, selector.Selector) model.FileInfos); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.FileInfos)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, selector.Selector) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileInfo_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type FileInfo_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 selector.Selector
func (_e *FileInfo_Expecter) GetAll(ctx interface{}, _a1 interface{}) *FileInfo_GetAll_Call {
	return &FileInfo_GetAll_Call{Call: _e.mock.On("GetAll", ctx, _a1)}
}

func (_c *FileInfo_GetAll_Call) Run(run func(ctx context.Context, _a1 selector.Selector)) *FileInfo_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(selector.Selector))
	})
	return _c
}

func (_c *FileInfo_GetAll_Call) Return(_a0 model.FileInfos, _a1 error) *FileInfo_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileInfo_GetAll_Call) RunAndReturn(run func(context.Context, selector.Selector) (model.FileInfos, error)) *FileInfo_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *FileInfo) GetByID(ctx context.Context, id string) (model.FileInfo, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FileInfo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FileInfo); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.FileInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileInfo_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type FileInfo_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *FileInfo_Expecter) GetByID(ctx interface{}, id interface{}) *FileInfo_GetByID_Call {
	return &FileInfo_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *FileInfo_GetByID_Call) Run(run func(ctx context.Context, id string)) *FileInfo_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileInfo_GetByID_Call) Return(_a0 model.FileInfo, _a1 error) *FileInfo_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileInfo_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.FileInfo, error)) *FileInfo_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileInfo creates a new instance of FileInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileInfo {
	mock := &FileInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}