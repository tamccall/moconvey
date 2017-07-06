package testify

import "github.com/stretchr/testify/mock"

// Mock is the assertion interface exposed by stretchr/testify. It exposes the various verification methods implemented
// by the mock struct
type Mock interface {
	AssertCalled(t mock.TestingT, methodName string, args... interface{}) bool
	AssertExpectations(t mock.TestingT) bool
	AssertNotCalled(t mock.TestingT, methodName string, args... interface{}) bool
	AssertNumberOfCalls(t mock.TestingT, methodName string, times int) bool
}
