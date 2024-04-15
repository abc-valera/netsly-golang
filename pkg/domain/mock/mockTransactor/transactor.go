// Code generated by mockery. DO NOT EDIT.

package mockTransactor

import (
	context "context"

	domain "github.com/abc-valera/netsly-api-golang/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// Transactor is an autogenerated mock type for the ITransactor type
type Transactor struct {
	mock.Mock
}

type Transactor_Expecter struct {
	mock *mock.Mock
}

func (_m *Transactor) EXPECT() *Transactor_Expecter {
	return &Transactor_Expecter{mock: &_m.Mock}
}

// PerformTX provides a mock function with given fields: ctx, txFunc
func (_m *Transactor) PerformTX(ctx context.Context, txFunc func(context.Context, domain.Entities) error) error {
	ret := _m.Called(ctx, txFunc)

	if len(ret) == 0 {
		panic("no return value specified for PerformTX")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, domain.Entities) error) error); ok {
		r0 = rf(ctx, txFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transactor_PerformTX_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PerformTX'
type Transactor_PerformTX_Call struct {
	*mock.Call
}

// PerformTX is a helper method to define mock.On call
//   - ctx context.Context
//   - txFunc func(context.Context , domain.Entities) error
func (_e *Transactor_Expecter) PerformTX(ctx interface{}, txFunc interface{}) *Transactor_PerformTX_Call {
	return &Transactor_PerformTX_Call{Call: _e.mock.On("PerformTX", ctx, txFunc)}
}

func (_c *Transactor_PerformTX_Call) Run(run func(ctx context.Context, txFunc func(context.Context, domain.Entities) error)) *Transactor_PerformTX_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context, domain.Entities) error))
	})
	return _c
}

func (_c *Transactor_PerformTX_Call) Return(_a0 error) *Transactor_PerformTX_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transactor_PerformTX_Call) RunAndReturn(run func(context.Context, func(context.Context, domain.Entities) error) error) *Transactor_PerformTX_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransactor creates a new instance of Transactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Transactor {
	mock := &Transactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
