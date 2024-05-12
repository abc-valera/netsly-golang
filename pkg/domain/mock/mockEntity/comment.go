// Code generated by mockery. DO NOT EDIT.

package mockEntity

import (
	context "context"

	entity "github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"

	selector "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
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

// Create provides a mock function with given fields: ctx, req
func (_m *Comment) Create(ctx context.Context, req entity.CommentCreateRequest) (model.Comment, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CommentCreateRequest) (model.Comment, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CommentCreateRequest) model.Comment); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CommentCreateRequest) error); ok {
		r1 = rf(ctx, req)
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
//   - req entity.CommentCreateRequest
func (_e *Comment_Expecter) Create(ctx interface{}, req interface{}) *Comment_Create_Call {
	return &Comment_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *Comment_Create_Call) Run(run func(ctx context.Context, req entity.CommentCreateRequest)) *Comment_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CommentCreateRequest))
	})
	return _c
}

func (_c *Comment_Create_Call) Return(_a0 model.Comment, _a1 error) *Comment_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_Create_Call) RunAndReturn(run func(context.Context, entity.CommentCreateRequest) (model.Comment, error)) *Comment_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, commentID
func (_m *Comment) Delete(ctx context.Context, commentID string) error {
	ret := _m.Called(ctx, commentID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, commentID)
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
//   - commentID string
func (_e *Comment_Expecter) Delete(ctx interface{}, commentID interface{}) *Comment_Delete_Call {
	return &Comment_Delete_Call{Call: _e.mock.On("Delete", ctx, commentID)}
}

func (_c *Comment_Delete_Call) Run(run func(ctx context.Context, commentID string)) *Comment_Delete_Call {
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

// GetAllByJokeID provides a mock function with given fields: ctx, jokeID, params
func (_m *Comment) GetAllByJokeID(ctx context.Context, jokeID string, params selector.Selector) (model.Comments, error) {
	ret := _m.Called(ctx, jokeID, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByJokeID")
	}

	var r0 model.Comments
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Comments, error)); ok {
		return rf(ctx, jokeID, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Comments); ok {
		r0 = rf(ctx, jokeID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Comments)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, jokeID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Comment_GetAllByJokeID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByJokeID'
type Comment_GetAllByJokeID_Call struct {
	*mock.Call
}

// GetAllByJokeID is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
//   - params selector.Selector
func (_e *Comment_Expecter) GetAllByJokeID(ctx interface{}, jokeID interface{}, params interface{}) *Comment_GetAllByJokeID_Call {
	return &Comment_GetAllByJokeID_Call{Call: _e.mock.On("GetAllByJokeID", ctx, jokeID, params)}
}

func (_c *Comment_GetAllByJokeID_Call) Run(run func(ctx context.Context, jokeID string, params selector.Selector)) *Comment_GetAllByJokeID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Comment_GetAllByJokeID_Call) Return(_a0 model.Comments, _a1 error) *Comment_GetAllByJokeID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_GetAllByJokeID_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Comments, error)) *Comment_GetAllByJokeID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Comment) GetByID(ctx context.Context, id string) (model.Comment, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Comment, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Comment); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Comment_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type Comment_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Comment_Expecter) GetByID(ctx interface{}, id interface{}) *Comment_GetByID_Call {
	return &Comment_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *Comment_GetByID_Call) Run(run func(ctx context.Context, id string)) *Comment_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Comment_GetByID_Call) Return(_a0 model.Comment, _a1 error) *Comment_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.Comment, error)) *Comment_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, commentID, req
func (_m *Comment) Update(ctx context.Context, commentID string, req entity.CommentUpdateRequest) (model.Comment, error) {
	ret := _m.Called(ctx, commentID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.CommentUpdateRequest) (model.Comment, error)); ok {
		return rf(ctx, commentID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.CommentUpdateRequest) model.Comment); ok {
		r0 = rf(ctx, commentID, req)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.CommentUpdateRequest) error); ok {
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
//   - req entity.CommentUpdateRequest
func (_e *Comment_Expecter) Update(ctx interface{}, commentID interface{}, req interface{}) *Comment_Update_Call {
	return &Comment_Update_Call{Call: _e.mock.On("Update", ctx, commentID, req)}
}

func (_c *Comment_Update_Call) Run(run func(ctx context.Context, commentID string, req entity.CommentUpdateRequest)) *Comment_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.CommentUpdateRequest))
	})
	return _c
}

func (_c *Comment_Update_Call) Return(_a0 model.Comment, _a1 error) *Comment_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Comment_Update_Call) RunAndReturn(run func(context.Context, string, entity.CommentUpdateRequest) (model.Comment, error)) *Comment_Update_Call {
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
