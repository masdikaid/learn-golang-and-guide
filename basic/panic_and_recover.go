package basic

import "fmt"

// panic is func that can build error scenario in code
// to make our app no stop with our own error we need catch it with recover using defer
func Panic() {
	defer errorHandler()
	a := true
	if a {
		panic("Waduh Error Lagi")
	}
	fmt.Println("aplikasi berjalan")
}

func errorHandler() {
	err := recover()
	if err != nil {
		fmt.Println("terjadi error ! ", err)
	}
	fmt.Println("logged ...")
}
