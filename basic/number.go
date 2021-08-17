package basic

import (
	"fmt"
	"reflect"
)

// golang has 2 type of number
// -integer : int8/uint8, int16/uint16, int32/uint32, int64/uint64 (default integer following your os bit 32/64)
// -float : float32, float64, complex64, complex128

//alias :
// byte = uint8
// rune = int32
// int = minimal int32 (follow your os)
// uint = minimal uint32 (follow your os)
func Number() {
	fmt.Println(232, reflect.TypeOf(232))
	fmt.Println(232.67, reflect.TypeOf(232.67))
}
