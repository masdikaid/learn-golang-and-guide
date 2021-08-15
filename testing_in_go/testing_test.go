package testingingo

// builtin package for testing
import (
	"fmt"
	"testing"
)

// file name for test must be end with '_test'
// naming test function must start with Test
// running test with go test in directory or go test -v for detail running list
// running test with go test ./... in root to run all test in directory and child directory
// add run parameter to run specific test ie. 'go test ./... -run TestMy -v' will run all test that contain 'TestMy'

func TestMyFuncWithFail(t *testing.T) {
	result := MyFunc()
	if result != "Mytest" {
		// t.Fail will stop running code and continue to next code line
		t.Fail()
	}
	// when using fail next line still executed
	fmt.Println(">>>> Printed <<<<")
}

func TestMyFuncWithFailNow(t *testing.T) {
	result := MyFunc()
	if result != "Mytest" {
		// t.FailNow will stop running test and continue to next test
		t.FailNow()
	}
	// when using fail now next line still skiped and runing next step
	fmt.Println(">>>> Printed <<<<")
}

func TestMyFuncWithError(t *testing.T) {
	result := MyFunc()
	if result != "Mytest" {
		// t.Error or t.Errorf is like Fail() But with message
		t.Errorf("result must be 'Mytest' but '%s' given", result)
	}
	// when using error next line still executed
	fmt.Println(">>>> Printed <<<<")
}

func TestMyFuncWithFatal(t *testing.T) {
	result := MyFunc()
	if result != "Mytest" {
		// t.Fatal or t.Fatalf is like FailNow() But with message
		t.Fatalf("result must be 'Mytest' but '%s' given", result)
	}
	// when using fatal next line still skiped and runing next step
	fmt.Println(">>>> Printed <<<<")
}
