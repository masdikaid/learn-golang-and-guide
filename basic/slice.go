package basic

import (
	"fmt"
	"reflect"
)

// slice is piece of array that represent the array
// when you change the value of the slice it will change the value in the array

func SliceData() {
	// array
	months := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	// create slice from existing array
	slice := months[2:9]
	// create new slice
	newslice := make([]int, 3, 10) //type, pointer, length
	// create new slice with [] without length of array like [3]
	newslice2 := []int{6, 7, 8, 9, 0}

	fmt.Println(slice, newslice, newslice2, reflect.TypeOf(newslice))
}
