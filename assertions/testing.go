package assertions

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
)

// fakeT is a struct intended to match the mock packages TestingT interface
// see http://bit.ly/2sqvoLE
type fakeT struct {
	buf    bytes.Buffer
	failed bool
}

func (t *fakeT) Logf(format string, args ...interface{}) {
	fmt.Fprintf(&t.buf, format, args...)
}

func (t *fakeT) Errorf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.FailNow()
}

func (t *fakeT) FailNow() {
	t.failed = true
}

func (t *fakeT) error() string {
	return t.buf.String()
}

var _ assert.TestingT = (*fakeT)(nil)
