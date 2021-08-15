package testingingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// golang support func test inside func test using testing.T.Run
// comman to subtest go test ./... -v -run TestMyFuncSubTest/Sub_Test_Name or go test ./... -v -run /Sub_Test_Name

func TestMyFuncSubTest(t *testing.T) {
	t.Run("Sub Test Name", func(t *testing.T) {
		assert.Equal(t, MyFunc(), "Mytest")
	})

	t.Run("Sub Second Test Name", func(t *testing.T) {
		assert.Equal(t, MyFunc(), "Mytest")
	})

	// Result :
	// === RUN   TestMyFuncSubTest
	// === RUN   TestMyFuncSubTest/Sub_Test_Name
	// === RUN   TestMyFuncSubTest/Sub_Second_Test_Name
	// --- PASS: TestMyFuncSubTest (0.00s)
	//    --- PASS: TestMyFuncSubTest/Sub_Test_Name (0.00s)
	//    --- PASS: TestMyFuncSubTest/Sub_Second_Test_Name (0.00s)
}
