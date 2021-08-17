package basic

import "fmt"

// operator in go is most likely other language
// https://www.tutorialspoint.com/go/go_operators.htm
func Operator() {
	a, b := 4, 2

	fmt.Println("+ for add 1 + 1 = ", a+b)
	fmt.Println("- for substract 1 - 1 = ", a-b)
	fmt.Println("/ for divide 4 : 2 = ", a/b)
	fmt.Println("* for multiplication 4 x 2 = ", a*b)
	fmt.Println("% for modulus 4 % 2 = ", a%b)
	a++
	fmt.Println("++ for add 1 = ", a)
	a--
	fmt.Println("-- for substract 1 = ", a)
}
