package testing_in_go

import (
	"fmt"
	"testing"
)

// basically noting before and after test in golang but there is func that can work like before and after test
// using testing.M (Main) we can setup before and after test
// Test Main just run once per package

// named func must be TestMain so golang can notice there is Main Function to run before runing all test and using testing.M in parameter
func TestMain(m *testing.M) {
	// write before code here
	fmt.Println("before testing")

	// run test
	m.Run()

	// write after code here
	fmt.Println("after testing")
}
