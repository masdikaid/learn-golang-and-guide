package basic

import "fmt"

func ForLoop() {
	// create for loop
	counter := 0
	for counter <= 10 {
		fmt.Print(counter)
		counter++
	}
	fmt.Print("\n")

	// create for loop and declare counter immediately
	for counter := 0; counter <= 10; counter++ {
		fmt.Print(counter)
	}
	fmt.Print("\n")

	// create for range from array
	array := []string{"dika", "adi", "setiawan"}
	for i, name := range array {
		fmt.Println(i, name)
	}

	// create for range from array without index
	array2 := []string{"dika", "adi", "setiawan"}
	for _, name := range array2 {
		fmt.Print(name, " ")
	}
	fmt.Print("\n")

}
