# moconvey
[![Build Status](https://travis-ci.org/nukosuke/go-zendesk.svg?branch=master)](https://travis-ci.org/nukosuke/go-zendesk)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamccall/moconvey)](https://goreportcard.com/report/github.com/tamccall/moconvey)
[![GoDoc](https://godoc.org/github.com/tamccall/moconvey?status.svg)](https://godoc.org/github.com/tamccall/moconvey)

moconvey is a package that aims to better integrate the [mocks](https://godoc.org/github.com/stretchr/testify/mock)
of [stretchr/testify](https://github.com/stretchr/testify) and the testing framework that is [goconvey](https://github.com/smartystreets/goconvey)

This package currently defines all of the assertions provided by the [mock struct](https://godoc.org/github.com/stretchr/testify/mock#Mock)
as [convey style assertions](https://github.com/smartystreets/goconvey/wiki/Custom-Assertions) so you will still be able
to get those fancy green check marks for your mock assertions too.

This package was originally designed to work with mocks generated by [mockery](https://github.com/vektra/mockery) but should
work with any mock using testify's mock package

## Example
Here is an example test function testing a piece of code that operates on a mock called `Foo`.
The test function can setup expectations (testify) for `Foo` and assert that they indeed happened:

```go
package convey

import (
	"testing"
	
	. "github.com/tamccall/moconvey/assertions"
	. "github.com/smartystreets/goconvey/convey"
	
	"github.com/tamccall/moconvey/mocks"	
)

func TestExampleMock(t *testing.T)	{
	Convey( "Given the example mock" , t, func() {
		mock := new(mocks.Foo)
		mock.On("Bar").Return("Hello World")
		Convey("When Bar is called", func() {
			mock.Bar()
			Convey("Assert Bar is called", func() {
				So(mock, ShouldHaveReceived, "Bar")
			})
		})
	})
}
```
