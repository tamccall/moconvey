// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Foo is an autogenerated mock type for the Foo type
type Foo struct {
	mock.Mock
}

// Bar provides a mock function with given fields:
func (_m *Foo) Bar() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
