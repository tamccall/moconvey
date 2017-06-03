package assertions

import (
	"fmt"
	"reflect"

	"github.com/stretchr/testify/mock"
)

const (
	testPassedMessage = ""
)

func getMock(iFace interface{}) (mock.Mock, error) {
	if m, ok := iFace.(mock.Mock); ok {
		return m, nil
	}

	var out mock.Mock
	iVal := reflect.ValueOf(iFace)
	if iVal.Type().Kind() == reflect.Ptr {
		return getMock(iVal.Elem().Interface())
	}

	if iVal.Type().Kind() == reflect.Struct && iVal.NumField() >= 1 {
		v := iVal.Field(0)
		if m, ok := v.Interface().(mock.Mock); ok {
			out = m
		} else {
			return mock.Mock{}, fmt.Errorf("Could not get mock from interface %v. Had type %v", iFace, reflect.TypeOf(iFace))
		}
	} else {
		return mock.Mock{}, fmt.Errorf("Could not get mock from interface %v. Had type %v", iFace, reflect.TypeOf(iFace))
	}

	return out, nil
}

// ShouldHaveReceived is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertCalled
// , and it asserts that the specified method was called.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Assertions
//
func ShouldHaveReceived(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}
	var methodName string
	if stringVal, ok := args[0].(string); !ok {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
	} else {
		methodName = stringVal
	}

	passed := m.AssertCalled(t, methodName, args[1:]...)

	if !passed {
		return t.error()
	}

	return testPassedMessage
}

// ShouldHaveReceivedN is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertNumberOfCalls
// , and it asserts that the specified method was called N times.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Assertions
func ShouldHaveReceivedN(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}

	if len(args) < 2 {
		return fmt.Sprint("Recived less that the expected 2 arguments")
	}

	var methodName string
	if stringVal, ok := args[0].(string); !ok {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
	} else {
		methodName = stringVal
	}

	var times int
	if intVal, ok := args[1].(int); ok {
		times = intVal
	} else {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 1 (times) should be an int. (was %s)", typ.Name())
	}

	passed := m.AssertNumberOfCalls(t, methodName, times)

	if !passed {
		return t.error()
	}

	return testPassedMessage
}

// ShouldNotHaveReceived is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertNotCalled
// , and it asserts that the specified method was not called.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Assertions
func ShouldNotHaveReceived(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}
	var methodName string
	if stringVal, ok := args[0].(string); !ok {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
	} else {
		methodName = stringVal
	}

	passed := m.AssertNotCalled(t, methodName, args[1:]...)

	if !passed {
		return t.error()
	}

	return testPassedMessage
}

// HadExpectationsMet is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertExpectations, and asserts that
// everything specified with On and Return was in fact called as expected.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Assertions
func HadExpectationsMet(iFace interface{}, args ...interface{}) string {
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}

	t := new(fakeT)
	passed := m.AssertExpectations(t)

	if !passed {
		return t.error()
	}

	return testPassedMessage
}
