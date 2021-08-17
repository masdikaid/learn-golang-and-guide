package basic

import "fmt"

// in go you can access variable in parent block of function from child function, so be patient when use it
func ClosureInFunc() {
	a := 1
	b := 2
	myFu := func() {
		// this will access variable in parrent block
		a = 5
		// this will create new variable just for myFu func
		b := 3
		fmt.Println(a)
		b++
		fmt.Println(b)
	}
	myFu()
	fmt.Println(a)
	fmt.Println(b)

}
