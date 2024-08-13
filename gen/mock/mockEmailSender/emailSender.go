// Code generated by mockery. DO NOT EDIT.

package mockEmailSender

import (
	service "github.com/abc-valera/netsly-golang/internal/domain/service"
	mock "github.com/stretchr/testify/mock"
)

// EmailSender is an autogenerated mock type for the IEmailSender type
type EmailSender struct {
	mock.Mock
}

type EmailSender_Expecter struct {
	mock *mock.Mock
}

func (_m *EmailSender) EXPECT() *EmailSender_Expecter {
	return &EmailSender_Expecter{mock: &_m.Mock}
}

// SendEmail provides a mock function with given fields: e
func (_m *EmailSender) SendEmail(e service.Email) error {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(service.Email) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmailSender_SendEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmail'
type EmailSender_SendEmail_Call struct {
	*mock.Call
}

// SendEmail is a helper method to define mock.On call
//   - e service.Email
func (_e *EmailSender_Expecter) SendEmail(e interface{}) *EmailSender_SendEmail_Call {
	return &EmailSender_SendEmail_Call{Call: _e.mock.On("SendEmail", e)}
}

func (_c *EmailSender_SendEmail_Call) Run(run func(e service.Email)) *EmailSender_SendEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(service.Email))
	})
	return _c
}

func (_c *EmailSender_SendEmail_Call) Return(_a0 error) *EmailSender_SendEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EmailSender_SendEmail_Call) RunAndReturn(run func(service.Email) error) *EmailSender_SendEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewEmailSender creates a new instance of EmailSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailSender(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailSender {
	mock := &EmailSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
