// Code generated by mockery. DO NOT EDIT.

package mockCommand

import (
	context "context"

	command "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"

	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

// Comment is an autogenerated mock type for the IComment type
type Comment struct {
	mock.Mock
}

type Comment_Expecter struct {
	mock *mock.Mock
}

func (_m *Comment) EXPECT() *Comment_Expecter {
	return &Comment_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, userID, jokeID, req
func (_m *Comment) Create(ctx context.Context, userID string, jokeID string, req model.Comment) (model.Comment, error) {
	ret := _m.Called(ctx, userID, jokeID, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.Comment) (model.Comment, error)); ok {
		return rf(ctx, userID, jokeID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.Comment) model.Comment); ok {
		r0 = rf(ctx, userID, jokeID, req)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, model.Comment) error); ok {
		r1 = rf(ctx, userID, jokeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Comment_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Comment_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - jokeID string
//   - req model.Comment
func (_e *Comment_Expecter) Create(ctx interface{}, userID interface{}, jokeID interface{}, req interface{}) *Comment_Create_Call {
	return &Comment_Create_Call{Call: _e.mock.On("Create", ctx, userID, jokeID, req)}
}

func (_c *Comment_Create_Call) Run(run func(ctx context.Context, userID string, jokeID string, req model.Comment)) *Comment_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(model.Comment))
	})
	return _c
}

func (_c *Comment_Create_Call) Return(_a0 model.Comment, _a1 error) *Comment_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_Create_Call) RunAndReturn(run func(context.Context, string, string, model.Comment) (model.Comment, error)) *Comment_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Comment) Delete(ctx context.Context, id string) error {
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

// Comment_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Comment_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Comment_Expecter) Delete(ctx interface{}, id interface{}) *Comment_Delete_Call {
	return &Comment_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *Comment_Delete_Call) Run(run func(ctx context.Context, id string)) *Comment_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Comment_Delete_Call) Return(_a0 error) *Comment_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Comment_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Comment_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, commentID, req
func (_m *Comment) Update(ctx context.Context, commentID string, req command.CommentUpdate) (model.Comment, error) {
	ret := _m.Called(ctx, commentID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, command.CommentUpdate) (model.Comment, error)); ok {
		return rf(ctx, commentID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, command.CommentUpdate) model.Comment); ok {
		r0 = rf(ctx, commentID, req)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, command.CommentUpdate) error); ok {
		r1 = rf(ctx, commentID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Comment_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Comment_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - commentID string
//   - req command.CommentUpdate
func (_e *Comment_Expecter) Update(ctx interface{}, commentID interface{}, req interface{}) *Comment_Update_Call {
	return &Comment_Update_Call{Call: _e.mock.On("Update", ctx, commentID, req)}
}

func (_c *Comment_Update_Call) Run(run func(ctx context.Context, commentID string, req command.CommentUpdate)) *Comment_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(command.CommentUpdate))
	})
	return _c
}

func (_c *Comment_Update_Call) Return(_a0 model.Comment, _a1 error) *Comment_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_Update_Call) RunAndReturn(run func(context.Context, string, command.CommentUpdate) (model.Comment, error)) *Comment_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewComment creates a new instance of Comment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComment(t interface {
	mock.TestingT
	Cleanup(func())
}) *Comment {
	mock := &Comment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
