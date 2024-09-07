// Code generated by mockery. DO NOT EDIT.

package mockEntity

import (
	entity "github.com/abc-valera/netsly-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// Emailer is an autogenerated mock type for the IEmailer type
type Emailer struct {
	mock.Mock
}

type Emailer_Expecter struct {
	mock *mock.Mock
}

func (_m *Emailer) EXPECT() *Emailer_Expecter {
	return &Emailer_Expecter{mock: &_m.Mock}
}

// SendEmail provides a mock function with given fields: e
func (_m *Emailer) SendEmail(e entity.EmailSendRequest) error {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.EmailSendRequest) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Emailer_SendEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmail'
type Emailer_SendEmail_Call struct {
	*mock.Call
}

// SendEmail is a helper method to define mock.On call
//   - e entity.EmailSendRequest
func (_e *Emailer_Expecter) SendEmail(e interface{}) *Emailer_SendEmail_Call {
	return &Emailer_SendEmail_Call{Call: _e.mock.On("SendEmail", e)}
}

func (_c *Emailer_SendEmail_Call) Run(run func(e entity.EmailSendRequest)) *Emailer_SendEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.EmailSendRequest))
	})
	return _c
}

func (_c *Emailer_SendEmail_Call) Return(_a0 error) *Emailer_SendEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Emailer_SendEmail_Call) RunAndReturn(run func(entity.EmailSendRequest) error) *Emailer_SendEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewEmailer creates a new instance of Emailer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Emailer {
	mock := &Emailer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
