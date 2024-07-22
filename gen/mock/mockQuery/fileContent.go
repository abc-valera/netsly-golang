// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// FileContent is an autogenerated mock type for the IFileContent type
type FileContent struct {
	mock.Mock
}

type FileContent_Expecter struct {
	mock *mock.Mock
}

func (_m *FileContent) EXPECT() *FileContent_Expecter {
	return &FileContent_Expecter{mock: &_m.Mock}
}

// GetByID provides a mock function with given fields: ctx, fileID
func (_m *FileContent) GetByID(ctx context.Context, fileID string) (model.FileContent, error) {
	ret := _m.Called(ctx, fileID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.FileContent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FileContent, error)); ok {
		return rf(ctx, fileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FileContent); ok {
		r0 = rf(ctx, fileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.FileContent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileContent_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type FileContent_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - fileID string
func (_e *FileContent_Expecter) GetByID(ctx interface{}, fileID interface{}) *FileContent_GetByID_Call {
	return &FileContent_GetByID_Call{Call: _e.mock.On("GetByID", ctx, fileID)}
}

func (_c *FileContent_GetByID_Call) Run(run func(ctx context.Context, fileID string)) *FileContent_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileContent_GetByID_Call) Return(_a0 model.FileContent, _a1 error) *FileContent_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileContent_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.FileContent, error)) *FileContent_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileContent creates a new instance of FileContent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileContent(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileContent {
	mock := &FileContent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
