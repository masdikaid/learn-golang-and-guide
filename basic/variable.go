package basic

import "fmt"

// variable is very helpfull and will very needed
// variable can save your data and you can access or change the value when you needed
// change value must be in the same type
// variable just can contain 1 type of data and you can't change it when it was created

func Variable() {
	// declare variable
	var myvar string
	// or you can direct filling your variable and the variable type will set automatically
	var my2var = true
	// or you can use := to declaring variable without var, but you have to fill your variable immediately
	my3var := 63
	// declare multi variable
	var (
		name = "joko"
		age  = 22
	)

	// fill your variable
	myvar = "fillfulled"

	// change value
	myvar = "new change"
	my2var = false
	my3var = 100

	// accessing your variable
	fmt.Println(myvar, my2var, my3var, name, age)
}
