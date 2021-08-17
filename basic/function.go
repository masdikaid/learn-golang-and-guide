package basic

import "fmt"

func Function() {
	// call function
	myFunc()
	myFuncWithParams("dika", "setiawan")
	// function can be value of variable if we not use ()
	paramsFunc := myFuncWithParams
	paramsFunc("dika", "Adi Setiawan")

	fmt.Println(myFuncWithReturn("dika"))
	fmt.Println(myFuncWithNamedReturn("yuda"))
	a, b, c := myFuncWithMultiReturn("dika")
	fmt.Println(a, b, c)

	// use underscore if just need one value of return multiple data
	d, _, _ := myFuncWithMultiReturn("dika")
	fmt.Println(d)
	myVariadicFunc("dika", "merah", "kuning", "hijau")

	// input variadic params with slice using variable argument ("...")
	colors := []string{"merah", "kuning", "biru", "jingga"}
	myVariadicFunc("duma", colors...)

	// function as parameter
	wordFilter("joko", myFuncWithReturn)

	// anonnymous function
	anonFunction := func(name string) string {
		return name
	}
	fmt.Println(anonFunction("masdikaid"))
}

// declare function
func myFunc() {
	// your other code
	fmt.Println("this is my first function")
}

// declare function with parameter
func myFuncWithParams(name string, lastname string) {
	// your other code
	fmt.Println("this is my function with params", name, lastname)
}

// declare function with return value
func myFuncWithReturn(name string) string {
	// your other code
	return name + " Setiawan"
}

// declare function with named return value
func myFuncWithNamedReturn(name string) (fullname string) {
	// your other code
	fullname = name + " adi setiawan"
	return
}

// declare function with return multi value
func myFuncWithMultiReturn(name string) (string, int, string) {
	// your other code
	return name, 2, name + " Setiawan"
}

// variadic function
// with variadic, last parameter can be varargs or multi params using "..."
func myVariadicFunc(name string, favcolor ...string) {
	fmt.Println("my name is ", name)
	fmt.Print("my fav colors is : ")
	for _, color := range favcolor {
		fmt.Print(color, " ")
	}
	fmt.Println()
}

// function as parameter
// thats use full if we need put a function in parameter or you can declare function as type declaration first if the params is to long
func wordFilter(word string, filter func(string) string) {
	wordFilter := filter(word)
	fmt.Println(wordFilter)
}
