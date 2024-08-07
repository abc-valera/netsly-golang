// Code generated by mockery. DO NOT EDIT.

package mockQuery

import (
	context "context"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	selector "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

// Like is an autogenerated mock type for the ILike type
type Like struct {
	mock.Mock
}

type Like_Expecter struct {
	mock *mock.Mock
}

func (_m *Like) EXPECT() *Like_Expecter {
	return &Like_Expecter{mock: &_m.Mock}
}

// CountByJokeID provides a mock function with given fields: ctx, jokeID
func (_m *Like) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	ret := _m.Called(ctx, jokeID)

	if len(ret) == 0 {
		panic("no return value specified for CountByJokeID")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, jokeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, jokeID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jokeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Like_CountByJokeID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountByJokeID'
type Like_CountByJokeID_Call struct {
	*mock.Call
}

// CountByJokeID is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
func (_e *Like_Expecter) CountByJokeID(ctx interface{}, jokeID interface{}) *Like_CountByJokeID_Call {
	return &Like_CountByJokeID_Call{Call: _e.mock.On("CountByJokeID", ctx, jokeID)}
}

func (_c *Like_CountByJokeID_Call) Run(run func(ctx context.Context, jokeID string)) *Like_CountByJokeID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Like_CountByJokeID_Call) Return(_a0 int, _a1 error) *Like_CountByJokeID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Like_CountByJokeID_Call) RunAndReturn(run func(context.Context, string) (int, error)) *Like_CountByJokeID_Call {
	_c.Call.Return(run)
	return _c
}

// GatAllByJokeID provides a mock function with given fields: ctx, jokeID, _a2
func (_m *Like) GatAllByJokeID(ctx context.Context, jokeID string, _a2 selector.Selector) (model.Likes, error) {
	ret := _m.Called(ctx, jokeID, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GatAllByJokeID")
	}

	var r0 model.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Likes, error)); ok {
		return rf(ctx, jokeID, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Likes); ok {
		r0 = rf(ctx, jokeID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, jokeID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Like_GatAllByJokeID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GatAllByJokeID'
type Like_GatAllByJokeID_Call struct {
	*mock.Call
}

// GatAllByJokeID is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
//   - _a2 selector.Selector
func (_e *Like_Expecter) GatAllByJokeID(ctx interface{}, jokeID interface{}, _a2 interface{}) *Like_GatAllByJokeID_Call {
	return &Like_GatAllByJokeID_Call{Call: _e.mock.On("GatAllByJokeID", ctx, jokeID, _a2)}
}

func (_c *Like_GatAllByJokeID_Call) Run(run func(ctx context.Context, jokeID string, _a2 selector.Selector)) *Like_GatAllByJokeID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Like_GatAllByJokeID_Call) Return(_a0 model.Likes, _a1 error) *Like_GatAllByJokeID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Like_GatAllByJokeID_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Likes, error)) *Like_GatAllByJokeID_Call {
	_c.Call.Return(run)
	return _c
}

// NewLike creates a new instance of Like. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLike(t interface {
	mock.TestingT
	Cleanup(func())
}) *Like {
	mock := &Like{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
