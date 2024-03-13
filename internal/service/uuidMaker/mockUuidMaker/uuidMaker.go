// Code generated by mockery. DO NOT EDIT.

package mockUuidMaker

import mock "github.com/stretchr/testify/mock"

// UuidMaker is an autogenerated mock type for the IUuidMaker type
type UuidMaker struct {
	mock.Mock
}

type UuidMaker_Expecter struct {
	mock *mock.Mock
}

func (_m *UuidMaker) EXPECT() *UuidMaker_Expecter {
	return &UuidMaker_Expecter{mock: &_m.Mock}
}

// NewUUID provides a mock function with given fields:
func (_m *UuidMaker) NewUUID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NewUUID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UuidMaker_NewUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewUUID'
type UuidMaker_NewUUID_Call struct {
	*mock.Call
}

// NewUUID is a helper method to define mock.On call
func (_e *UuidMaker_Expecter) NewUUID() *UuidMaker_NewUUID_Call {
	return &UuidMaker_NewUUID_Call{Call: _e.mock.On("NewUUID")}
}

func (_c *UuidMaker_NewUUID_Call) Run(run func()) *UuidMaker_NewUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UuidMaker_NewUUID_Call) Return(_a0 string) *UuidMaker_NewUUID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UuidMaker_NewUUID_Call) RunAndReturn(run func() string) *UuidMaker_NewUUID_Call {
	_c.Call.Return(run)
	return _c
}

// NewUuidMaker creates a new instance of UuidMaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUuidMaker(t interface {
	mock.TestingT
	Cleanup(func())
}) *UuidMaker {
	mock := &UuidMaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
