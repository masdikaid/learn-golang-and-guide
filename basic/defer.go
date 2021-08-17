package basic

import "fmt"

// defer function is func that will run in last of block code in function
// defer will still run even when error occured
func Defer() {
	defer loging()
	fmt.Println("welcome to My App")
}

func loging() {
	fmt.Println("Logged ...")
}
