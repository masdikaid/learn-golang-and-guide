package basic

import "fmt"

// break can use for stop the looping
func Break() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("found", i)
			break
		}
		fmt.Print(i, " ")
	}

}

// continue can use for skip the looping
func Continue() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Print("found ", i, " ")
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

}
