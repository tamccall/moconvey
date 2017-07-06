package assertions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tamccall/moconvey/mocks"
)

var t = new(testing.T)

func ExampleHadExpectationsMet() {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it", func() {
				mock.Bar()
				Convey("Its expectations are met", func() {
					So(mock, ShouldHaveExpectationsMet)
				})
			})
		})
	})
}

func TestHadExpectationsMet(t *testing.T) {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it", func() {
				mock.Bar()
				Convey("Its expectations are met", func() {
					So(mock, ShouldHaveExpectationsMet)
				})
			})
		})
	})
}

func ExampleShouldHaveReceived() {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it", func() {
				mock.Bar()
				Convey("It should have received the call", func() {
					So(mock, ShouldHaveReceived, "Bar")
				})
			})
		})
	})
}

func TestShouldHaveReceived(t *testing.T) {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it", func() {
				mock.Bar()
				Convey("It should have received the call", func() {
					So(mock, ShouldHaveReceived, "Bar")
				})
			})
		})
	})
}

func ExampleShouldHaveReceivedN() {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it twice", func() {
				mock.Bar()
				mock.Bar()
				Convey("We can assert that it was called twice", func() {
					So(mock, ShouldHaveReceivedN, "Bar", 2)
				})
			})
		})
	})
}

func TestShouldHaveReceivedN(t *testing.T) {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("After we call it twice", func() {
				mock.Bar()
				mock.Bar()
				Convey("We can assert that it was called twice", func() {
					So(mock, ShouldHaveReceivedN, "Bar", 2)
				})

				Convey("But it not 3 times", func() {
					msg := ShouldHaveReceived(mock, "Bar", 3)
					So(msg, ShouldNotEqual, testPassedMessage)
				})
			})
		})
	})
}

func ExampleShouldNotHaveReceived() {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("Can assert that bar hasn't been called", func() {
				So(mock, ShouldNotHaveReceived, "Bar")
			})

		})
	})
}

func TestShouldNotHaveReceived(t *testing.T) {
	Convey("Given A mock", t, func() {
		mock := new(mocks.Foo)
		Convey("That is expecting Bar to be called", func() {
			mock.On("Bar").Return("Hello World")
			Convey("Can assert that bar hasn't been called", func() {
				So(mock, ShouldNotHaveReceived, "Bar")
			})

		})
	})
}
