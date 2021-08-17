package basic

import "fmt"

func BuiltinFunction() {
	// len for get length of an object (string, array, map, etc)
	fmt.Println("length :", len("saya dika"))                // 5
	fmt.Println("length :", len([]int{1, 2, 3, 4, 5, 6, 7})) // 7

	// convert data type
	fmt.Println("convert int64 to int8 :", int8(int64(89)))
	fmt.Println("convert int to string :", string(rune(89)))
}
