package basic

import "fmt"

func IfExpression() {
	a := 2
	// declare if expression
	if a > 1 {
		fmt.Println("thats true")
	}

	// declare if else expression
	if a > 1 {
		fmt.Println("thats true")
	} else {
		fmt.Println("thats false")
	}

	b := 3
	// declare if else expression
	if b > 1 {
		fmt.Println("thats true")
	} else if b < 1 {
		fmt.Println("thats wrong")
	} else {
		fmt.Println("thats false")
	}

	// if short statement
	// we can declare variable directly when declare if
	if c := a + b; c > 2 {
		fmt.Println("thats right")
	}

}
