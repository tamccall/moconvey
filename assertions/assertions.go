package assertions

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/tamccall/moconvey/testify"
)

const (
	testPassedMessage = ""
)

func getMock(iFace interface{}) (testify.Mock, error) {
	if m, ok := iFace.(testify.Mock); !ok {
		return nil, errors.New("Cannot get mock")
	} else {
		return m, nil
	}
}

// ShouldHaveReceived is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertCalled
// , and it asserts that the specified method was called.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Custom-Assertions
func ShouldHaveReceived(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}
	var methodName string
	if stringVal, ok := args[0].(string); ok {
		methodName = stringVal
	} else {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
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
// https://github.com/smartystreets/goconvey/wiki/Custom-Assertions
func ShouldHaveReceivedN(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}

	if len(args) < 2 {
		return fmt.Sprint("Received less that the expected 2 arguments")
	}

	var methodName string
	if stringVal, ok := args[0].(string); ok {
		methodName = stringVal
	} else {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
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
// https://github.com/smartystreets/goconvey/wiki/Custom-Assertions
func ShouldNotHaveReceived(iFace interface{}, args ...interface{}) string {
	t := new(fakeT)
	m, err := getMock(iFace)
	if err != nil {
		return err.Error()
	}
	var methodName string
	if stringVal, ok := args[0].(string); ok {
		methodName = stringVal
	} else {
		typ := reflect.TypeOf(iFace)
		return fmt.Sprintf("Argument 0 (methodName) should be a string. (was %s)", typ.Name())
	}

	passed := m.AssertNotCalled(t, methodName, args[1:]...)

	if !passed {
		return t.error()
	}

	return testPassedMessage
}

// ShouldHaveExpectationsMet is a goconvey style assertion.
// It is similar to https://godoc.org/github.com/stretchr/testify/mock#Mock.AssertExpectations, and asserts that
// everything specified with On and Return was in fact called as expected.
//
// See Also
//
// https://github.com/smartystreets/goconvey/wiki/Custom-Assertions
func ShouldHaveExpectationsMet(iFace interface{}, args ...interface{}) string {
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
