package basic

import "fmt"

// arrays is collection of data that have the same type
func ArrayData() {
	//create array
	var array [4]int
	// create array with value immediately
	var array2 = [3]int{
		1,
		2,
		3,
	}

	// set or change value of array
	array[0] = 4
	array[3] = 2

	// access array
	fmt.Println(len(array))
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array[0])
}
