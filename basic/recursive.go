package basic

import "fmt"

// Go support recursive thats mean you can call function it self inside
func RecursiveFunc(n int) {
	if n > 1 {
		fmt.Print(n, ", ")
		RecursiveFunc(n - 1)
	} else {
		fmt.Println("1, selesai !")
	}
}
