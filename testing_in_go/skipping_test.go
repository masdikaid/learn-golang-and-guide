package testing_in_go

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// use t.skip to skiping test ie. when test in development or for specific OS etc
func TestSkipMyFunc(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("this test in development in MAC")
	}
	assert.Equal(t, MyFunc(), "Mytest")
}
