package basic

import "fmt"

// constant is mostly like variable but the value is constanta or can't be changed
func Constanta() {
	// declare constanta variable is using const syntax and fill the value immediately
	const myconst = "the value"
	// multiple constant declaration
	const (
		my2const = 1231
		my3const = true
	)

	// using constanta
	fmt.Println(myconst, my2const, my3const)
}
