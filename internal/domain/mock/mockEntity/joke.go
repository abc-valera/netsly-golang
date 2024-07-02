// Code generated by mockery. DO NOT EDIT.

package mockEntity

import (
	context "context"

	entity "github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-api-golang/internal/domain/model"

	selector "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

// Joke is an autogenerated mock type for the IJoke type
type Joke struct {
	mock.Mock
}

type Joke_Expecter struct {
	mock *mock.Mock
}

func (_m *Joke) EXPECT() *Joke_Expecter {
	return &Joke_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *Joke) Create(ctx context.Context, req entity.JokeCreateRequest) (model.Joke, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Joke
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.JokeCreateRequest) (model.Joke, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.JokeCreateRequest) model.Joke); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.Joke)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.JokeCreateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Joke_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Joke_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.JokeCreateRequest
func (_e *Joke_Expecter) Create(ctx interface{}, req interface{}) *Joke_Create_Call {
	return &Joke_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *Joke_Create_Call) Run(run func(ctx context.Context, req entity.JokeCreateRequest)) *Joke_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.JokeCreateRequest))
	})
	return _c
}

func (_c *Joke_Create_Call) Return(_a0 model.Joke, _a1 error) *Joke_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Joke_Create_Call) RunAndReturn(run func(context.Context, entity.JokeCreateRequest) (model.Joke, error)) *Joke_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, jokeID
func (_m *Joke) Delete(ctx context.Context, jokeID string) error {
	ret := _m.Called(ctx, jokeID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, jokeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Joke_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Joke_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
func (_e *Joke_Expecter) Delete(ctx interface{}, jokeID interface{}) *Joke_Delete_Call {
	return &Joke_Delete_Call{Call: _e.mock.On("Delete", ctx, jokeID)}
}

func (_c *Joke_Delete_Call) Run(run func(ctx context.Context, jokeID string)) *Joke_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Joke_Delete_Call) Return(_a0 error) *Joke_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Joke_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Joke_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByUserID provides a mock function with given fields: ctx, userID, params
func (_m *Joke) GetAllByUserID(ctx context.Context, userID string, params selector.Selector) (model.Jokes, error) {
	ret := _m.Called(ctx, userID, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByUserID")
	}

	var r0 model.Jokes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Jokes, error)); ok {
		return rf(ctx, userID, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Jokes); ok {
		r0 = rf(ctx, userID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Jokes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Joke_GetAllByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUserID'
type Joke_GetAllByUserID_Call struct {
	*mock.Call
}

// GetAllByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - params selector.Selector
func (_e *Joke_Expecter) GetAllByUserID(ctx interface{}, userID interface{}, params interface{}) *Joke_GetAllByUserID_Call {
	return &Joke_GetAllByUserID_Call{Call: _e.mock.On("GetAllByUserID", ctx, userID, params)}
}

func (_c *Joke_GetAllByUserID_Call) Run(run func(ctx context.Context, userID string, params selector.Selector)) *Joke_GetAllByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Joke_GetAllByUserID_Call) Return(_a0 model.Jokes, _a1 error) *Joke_GetAllByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Joke_GetAllByUserID_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Jokes, error)) *Joke_GetAllByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Joke) GetByID(ctx context.Context, id string) (model.Joke, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 model.Joke
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Joke, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Joke); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Joke)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Joke_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type Joke_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Joke_Expecter) GetByID(ctx interface{}, id interface{}) *Joke_GetByID_Call {
	return &Joke_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *Joke_GetByID_Call) Run(run func(ctx context.Context, id string)) *Joke_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Joke_GetByID_Call) Return(_a0 model.Joke, _a1 error) *Joke_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Joke_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.Joke, error)) *Joke_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAllByTitle provides a mock function with given fields: ctx, keyword, params
func (_m *Joke) SearchAllByTitle(ctx context.Context, keyword string, params selector.Selector) (model.Jokes, error) {
	ret := _m.Called(ctx, keyword, params)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByTitle")
	}

	var r0 model.Jokes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) (model.Jokes, error)); ok {
		return rf(ctx, keyword, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, selector.Selector) model.Jokes); ok {
		r0 = rf(ctx, keyword, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Jokes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, selector.Selector) error); ok {
		r1 = rf(ctx, keyword, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Joke_SearchAllByTitle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAllByTitle'
type Joke_SearchAllByTitle_Call struct {
	*mock.Call
}

// SearchAllByTitle is a helper method to define mock.On call
//   - ctx context.Context
//   - keyword string
//   - params selector.Selector
func (_e *Joke_Expecter) SearchAllByTitle(ctx interface{}, keyword interface{}, params interface{}) *Joke_SearchAllByTitle_Call {
	return &Joke_SearchAllByTitle_Call{Call: _e.mock.On("SearchAllByTitle", ctx, keyword, params)}
}

func (_c *Joke_SearchAllByTitle_Call) Run(run func(ctx context.Context, keyword string, params selector.Selector)) *Joke_SearchAllByTitle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(selector.Selector))
	})
	return _c
}

func (_c *Joke_SearchAllByTitle_Call) Return(_a0 model.Jokes, _a1 error) *Joke_SearchAllByTitle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Joke_SearchAllByTitle_Call) RunAndReturn(run func(context.Context, string, selector.Selector) (model.Jokes, error)) *Joke_SearchAllByTitle_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, jokeID, req
func (_m *Joke) Update(ctx context.Context, jokeID string, req entity.JokeUpdateRequest) (model.Joke, error) {
	ret := _m.Called(ctx, jokeID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Joke
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.JokeUpdateRequest) (model.Joke, error)); ok {
		return rf(ctx, jokeID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.JokeUpdateRequest) model.Joke); ok {
		r0 = rf(ctx, jokeID, req)
	} else {
		r0 = ret.Get(0).(model.Joke)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.JokeUpdateRequest) error); ok {
		r1 = rf(ctx, jokeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Joke_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Joke_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - jokeID string
//   - req entity.JokeUpdateRequest
func (_e *Joke_Expecter) Update(ctx interface{}, jokeID interface{}, req interface{}) *Joke_Update_Call {
	return &Joke_Update_Call{Call: _e.mock.On("Update", ctx, jokeID, req)}
}

func (_c *Joke_Update_Call) Run(run func(ctx context.Context, jokeID string, req entity.JokeUpdateRequest)) *Joke_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.JokeUpdateRequest))
	})
	return _c
}

func (_c *Joke_Update_Call) Return(_a0 model.Joke, _a1 error) *Joke_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Joke_Update_Call) RunAndReturn(run func(context.Context, string, entity.JokeUpdateRequest) (model.Joke, error)) *Joke_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewJoke creates a new instance of Joke. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJoke(t interface {
	mock.TestingT
	Cleanup(func())
}) *Joke {
	mock := &Joke{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
