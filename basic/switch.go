package basic

import "fmt"

func SwitchExpression() {

	name := "dika"
	// create switch
	switch name {
	case "andi":
		fmt.Println(name)
	case "rudi":
		fmt.Println(name)
	case "dika":
		fmt.Println(name)
	default:
		fmt.Println("nothing")
	}

	// create switch
	switch length := len(name); length {
	case 1:
		fmt.Println("to short")
	case 2:
		fmt.Println("good")
	case 3:
		fmt.Println("still good")
	default:
		fmt.Println("to long")
	}

	// create switch without parameter
	length := len(name)
	switch {
	case length < 2:
		fmt.Println("to short")
	case length > 2:
		fmt.Println("good")
	case length > 4:
		fmt.Println("still good")
	default:
		fmt.Println("to long")
	}
}
