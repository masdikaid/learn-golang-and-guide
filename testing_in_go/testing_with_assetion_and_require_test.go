package testing_in_go

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// for assertion and require test in golang need external library there is several library and one of them named testify
// command to install 'go get github.com/stretchr/testify'
// differential of assert and require is assert will call Fail() and require will call FailNow() when failed
func TestMyFuncWithAssert(t *testing.T) {
	assert.Equal(t, MyFunc(), "Mytest", "result must be 'Mytest'")
	// when using assert next line still executed
	fmt.Println(">>>> Printed <<<<")
}

func TestMyFuncWithRequire(t *testing.T) {
	require.Equal(t, MyFunc(), "Mytest", "result must be 'Mytest'")
	// when using require next line still skiped and runing next step
	fmt.Println(">>>> Printed <<<<")

}
