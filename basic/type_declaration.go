package basic

import "fmt"

// in go we can create our own type
func TypeDeclaration() {
	// create our custom type
	type Gender string
	type Marriage bool
	type Filter func(string) string

	// declaring variable for our type
	var dika Gender
	var status Marriage
	var filter Filter

	dika = "Male"
	status = true
	filter = myFuncWithReturn

	fmt.Println(dika, status, filter("dika"))
}
