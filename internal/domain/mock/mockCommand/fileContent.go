// Code generated by mockery. DO NOT EDIT.

package mockCommand

import (
	context "context"

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

// Create provides a mock function with given fields: ctx, fileID, content
func (_m *FileContent) Create(ctx context.Context, fileID string, content []byte) error {
	ret := _m.Called(ctx, fileID, content)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, fileID, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileContent_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type FileContent_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - fileID string
//   - content []byte
func (_e *FileContent_Expecter) Create(ctx interface{}, fileID interface{}, content interface{}) *FileContent_Create_Call {
	return &FileContent_Create_Call{Call: _e.mock.On("Create", ctx, fileID, content)}
}

func (_c *FileContent_Create_Call) Run(run func(ctx context.Context, fileID string, content []byte)) *FileContent_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *FileContent_Create_Call) Return(_a0 error) *FileContent_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileContent_Create_Call) RunAndReturn(run func(context.Context, string, []byte) error) *FileContent_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, fileID
func (_m *FileContent) Delete(ctx context.Context, fileID string) error {
	ret := _m.Called(ctx, fileID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, fileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileContent_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FileContent_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - fileID string
func (_e *FileContent_Expecter) Delete(ctx interface{}, fileID interface{}) *FileContent_Delete_Call {
	return &FileContent_Delete_Call{Call: _e.mock.On("Delete", ctx, fileID)}
}

func (_c *FileContent_Delete_Call) Run(run func(ctx context.Context, fileID string)) *FileContent_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileContent_Delete_Call) Return(_a0 error) *FileContent_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileContent_Delete_Call) RunAndReturn(run func(context.Context, string) error) *FileContent_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, fileID, newContent
func (_m *FileContent) Update(ctx context.Context, fileID string, newContent []byte) error {
	ret := _m.Called(ctx, fileID, newContent)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, fileID, newContent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileContent_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type FileContent_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - fileID string
//   - newContent []byte
func (_e *FileContent_Expecter) Update(ctx interface{}, fileID interface{}, newContent interface{}) *FileContent_Update_Call {
	return &FileContent_Update_Call{Call: _e.mock.On("Update", ctx, fileID, newContent)}
}

func (_c *FileContent_Update_Call) Run(run func(ctx context.Context, fileID string, newContent []byte)) *FileContent_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *FileContent_Update_Call) Return(_a0 error) *FileContent_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileContent_Update_Call) RunAndReturn(run func(context.Context, string, []byte) error) *FileContent_Update_Call {
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
