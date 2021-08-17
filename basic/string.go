package basic

import (
	"fmt"
	"reflect"
)

// string is type of text or caracter
// string must write inside ("")
// you can use [] for get substring ie. "mystring"[3]
func String() {
	fmt.Println("this is my text", reflect.TypeOf(""))
}
